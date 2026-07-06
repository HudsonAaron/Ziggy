package gactor

import (
	"container/heap"
	"main/deps/glog"
	"main/deps/gutil"
	"time"
)

// HandleInit回调函数
func (ga *gActorStatus[S]) handleInit(actor *GActor[S], handle HandleInit[S]) error {
	if handle == nil {
		return nil
	}
	state, err := handle(actor, ga.state)
	if err != nil {
		return err
	}
	ga.state = state
	return nil
}

// HandleCall回调函数
func (ga *gActorStatus[S]) handleCall(actor *GActor[S], gam *GActorMsg[S]) {
	if gam.handle == nil {
		return
	}
	reply, state, err := gam.handle(actor, gam.msg, ga.state)
	if err != nil {
		glog.Error("HandleCall err: %v", err)
		return
	}
	ga.state = state
	if gam.reply != nil {
		gam.reply.send(reply)
	}
}

// HandleInfo回调函数
func (ga *gActorStatus[S]) handleInfo(actor *GActor[S], gam *GActorMsg[S]) {
	if gam.handle == nil {
		return
	}
	_, state, err := gam.handle(actor, gam.msg, ga.state)
	if err != nil {
		glog.Error("HandleInfo err: %v", err)
		return
	}
	ga.state = state
}

// Terminate回调函数
func (ga *gActorStatus[S]) terminate(actor *GActor[S], gam *GActorMsg[S]) {
	if gam.handle == nil {
		if gam.reply != nil {
			gam.reply.send(nil)
		}
		return
	}
	reply, state, err := gam.handle(actor, gam.msg, ga.state)
	if err != nil {
		glog.Error("Terminate err: %v", err)
		return
	}
	ga.state = state
	if gam.reply != nil {
		gam.reply.send(reply)
	}
}

// 添加timer
func (ga *gActorStatus[S]) addTimer(gam *GActorMsg[S]) {
	gTimer := gam.msg.(*gActorTimer[S])
	ga.timerMu.Lock()
	heap.Push(&ga.timers, gTimer)
	ga.timerMu.Unlock()
	select {
	case ga.timerWakeup <- struct{}{}:
	default:
	}
}

// timer 过期回调函数 — 直接投递消息到状态机的消息队列
func (t *gActorTimer[S]) fire(ga *gActorStatus[S]) {
	gam := GActorMsg[S]{
		msgType: INFO,
		msg:     t.msg,
		handle:  t.handle,
	}
	ga.msgChan <- gam
}

// 取消timer
func (ga *gActorStatus[S]) cancelTimer(gam *GActorMsg[S]) {
	timerId := gam.msg.(string)
	ga.timerMu.Lock()
	for i, t := range ga.timers {
		if t.id == timerId {
			ga.timers[i] = ga.timers[len(ga.timers)-1]
			ga.timers = ga.timers[:len(ga.timers)-1]
			heap.Init(&ga.timers)
			ga.timerMu.Unlock()
			return
		}
	}
	ga.timerMu.Unlock()
}

// 启动timer调度器 — 单goroutine调度所有timer
// 注意：与 loop() goroutine 共享 ga.timers，所有堆操作必须持有 timerMu。
// 阻塞等待（time.After）前释放锁，唤醒后重新验证 timer 未被取消。
func (ga *gActorStatus[S]) startTimerScheduler() {
	for {
		ga.timerMu.Lock()
		if len(ga.timers) == 0 {
			ga.timerMu.Unlock()
			select {
			case <-ga.timerWakeup:
				continue
			case <-ga.timerDone:
				return
			}
		}

		next := ga.timers[0]
		wait := next.expireAt - gutil.TimestampMilli()

		if wait <= 0 {
			_ = heap.Pop(&ga.timers)
			ga.timerMu.Unlock()
			next.fire(ga)
			continue
		}

		// 记录等待的 timer ID，释放锁后阻塞等待
		waitId := next.id
		ga.timerMu.Unlock()

		select {
		case <-time.After(time.Duration(wait) * time.Millisecond):
			ga.timerMu.Lock()
			// 重新验证：timer 可能已被 cancelTimer 移除
			if len(ga.timers) > 0 && ga.timers[0].id == waitId && ga.timers[0].expireAt <= gutil.TimestampMilli() {
				_ = heap.Pop(&ga.timers)
				ga.timerMu.Unlock()
				next.fire(ga)
			} else {
				ga.timerMu.Unlock()
			}
		case <-ga.timerWakeup:
			continue
		case <-ga.timerDone:
			return
		}
	}
}
