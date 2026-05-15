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
	heap.Push(&ga.timers, gTimer)
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
	for i, t := range ga.timers {
		if t.id == timerId {
			ga.timers[i] = ga.timers[len(ga.timers)-1]
			ga.timers = ga.timers[:len(ga.timers)-1]
			heap.Init(&ga.timers)
			return
		}
	}
}

// 启动timer调度器 — 单goroutine调度所有timer
func (ga *gActorStatus[S]) startTimerScheduler() {
	for {
		if len(ga.timers) == 0 {
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
			next.fire(ga)
			continue
		}

		select {
		case <-time.After(time.Duration(wait) * time.Millisecond):
			_ = heap.Pop(&ga.timers)
			next.fire(ga)
		case <-ga.timerWakeup:
			continue
		case <-ga.timerDone:
			return
		}
	}
}
