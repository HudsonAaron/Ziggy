package gactor

import (
	"main/deps/glog"
	"main/deps/gutil"
	"time"
)

// HandleInit回调函数
func (ga *gActorStatus) handleInit(handle func(...interface{}) (GActorState, error)) error {
	if handle == nil {
		return nil
	}
	state, err := handle(ga.state)
	if err != nil {
		return err
	}
	ga.resetState(state)
	return nil
}

// HandleCall回调函数
func (ga *gActorStatus) handleCall(gam *GActorMsg) {
	if gam.handle == nil {
		return
	}
	// 传递state的副本，避免回调函数直接修改原始state
	reply, state, err := gam.handle(gam.msg, ga.state)
	if err != nil {
		glog.Error("HandleCall err: %v", err)
		return
	}
	ga.resetState(state)
	if gam.reply != nil {
		safeReply(gam.reply, reply)
	}
}

// HandleInfo回调函数
func (ga *gActorStatus) handleInfo(gam *GActorMsg) {
	if gam.handle == nil {
		return
	}
	// 传递state的副本，避免回调函数直接修改原始state
	_, state, err := gam.handle(gam.msg, ga.state)
	if err != nil {
		glog.Error("HandleInfo err: %v", err)
		return
	}
	ga.resetState(state)
}

// Terminate回调函数
func (ga *gActorStatus) terminate(gam *GActorMsg) {
	if gam.handle == nil {
		safeReply(gam.reply, nil)
		return
	}
	// 传递state的副本，避免回调函数直接修改原始state
	reply, state, err := gam.handle(gam.msg, ga.state)
	if err != nil {
		glog.Error("Terminate err: %v", err)
		return
	}
	ga.resetState(state)
	if gam.reply != nil {
		safeReply(gam.reply, reply)
	}
}

// 重置state
func (ga *gActorStatus) resetState(state interface{}) {
	switch state := any(state).(type) {
	case GActorState:
		ga.state = state
		return
	default:
		return
	}
}

// 添加timer
func (ga *gActorStatus) addTimer(gam *GActorMsg) {
	gTimer := gam.msg.(*gActorTimer)
	ga.timerMap.Store(gTimer.id, gTimer)
	// 启动timer
	go gTimer.timerCallback(&GActor{key: ga.key})
}

// timer 过期回调函数
func (t *gActorTimer) timerCallback(ga *GActor) {
	for {
		nowTime := gutil.TimestampMilli()
		select {
		case <-t.stopC:
			return
		case <-time.After(time.Duration(t.expireAt-nowTime) * time.Millisecond):
			ga.Info(t.msg, t.handle)
			ga.CancelTimer(t.id)
			return
		}
	}
}

// 取消timer
func (ga *gActorStatus) cancelTimer(gam *GActorMsg) {
	timerId := gam.msg.(string)
	ga.timerMap.Range(func(key, value interface{}) bool {
		keyStr := key.(string)
		if keyStr != timerId {
			return true
		}
		timer := value.(*gActorTimer)
		if timer.id == timerId {
			if timer.stopC != nil {
				close(timer.stopC)
			}
			// 移除timer
			ga.timerMap.Delete(key)
			return false
		}
		return true
	})
}

// 安全地向reply channel发送数据，避免向已关闭的channel发送数据导致panic
func safeReply(reply chan interface{}, data interface{}) {
	select {
	case reply <- data:
	case <-time.After(100 * time.Millisecond): // 防止永久阻塞
		// 如果发送超时，说明接收方可能已经不再等待，可以忽略
	}
}
