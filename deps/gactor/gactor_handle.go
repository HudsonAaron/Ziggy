package gactor

import (
	"main/deps/glog"
)

// HandleInit回调函数
func (ga *gActorStatus) HandleInit(handle func(...interface{}) (GActorState, error)) error {
	if handle == nil {
		return nil
	}
	state, err := handle(ga.state)
	if err != nil {
		return err
	}
	ga.ResetState(state)
	return nil
}

// HandleCall回调函数
func (ga *gActorStatus) HandleCall(gam *GActorMsg) {
	if gam.handle == nil {
		return
	}
	// 传递state的副本，避免回调函数直接修改原始state
	reply, state, err := gam.handle(gam.msg, ga.state)
	if err != nil {
		glog.Error("HandleCall err: %v", err)
		return
	}
	ga.ResetState(state)
	if gam.reply != nil {
		gam.reply <- reply
	}
}

// HandleInfo回调函数
func (ga *gActorStatus) HandleInfo(gam *GActorMsg) {
	if gam.handle == nil {
		return
	}
	// 传递state的副本，避免回调函数直接修改原始state
	_, state, err := gam.handle(gam.msg, ga.state)
	if err != nil {
		glog.Error("HandleInfo err: %v", err)
		return
	}
	ga.ResetState(state)
}

// Terminate回调函数
func (ga *gActorStatus) Terminate(gam *GActorMsg) {
	if gam.handle == nil {
		gam.reply <- nil
		return
	}
	// 传递state的副本，避免回调函数直接修改原始state
	reply, state, err := gam.handle(gam.msg, ga.state)
	if err != nil {
		glog.Error("Terminate err: %v", err)
		return
	}
	ga.ResetState(state)
	if gam.reply != nil {
		gam.reply <- reply
	}
}

// 重置state
func (ga *gActorStatus) ResetState(state interface{}) {
	switch state := any(state).(type) {
	case GActorState:
		ga.state = state
		return
	default:
		return
	}
}
