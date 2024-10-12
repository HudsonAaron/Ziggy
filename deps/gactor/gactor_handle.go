package gactor

import (
	"main/deps/glog"
)

// HandleInit回调函数
func (ga *GActor) HandleInit(handle func(...interface{}) (GActorState, error)) error {
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
func (ga *GActor) HandleCall(gam *GActorMsg) {
	if gam.handle == nil {
		return
	}
	reply, state, err := gam.handle(gam.msg, ga.state)
	if err != nil {
		glog.Error(err.Error())
		return
	}
	ga.ResetState(state)
	gam.reply <- reply
}

// HandleInfo回调函数
func (ga *GActor) HandleInfo(gam *GActorMsg) {
	if gam.handle == nil {
		return
	}
	_, state, err := gam.handle(gam.msg, ga.state)
	if err != nil {
		glog.Error(err.Error())
		return
	}
	ga.ResetState(state)
}

// Terminate回调函数
func (ga *GActor) Terminate(gam *GActorMsg) {
	if gam.handle == nil {
		return
	}
	_, state, err := gam.handle(gam.msg, ga.state)
	if err != nil {
		glog.Error(err.Error())
		return
	}
	ga.ResetState(state)
}

// 重置state
func (ga *GActor) ResetState(state interface{}) {
	switch state := any(state).(type) {
	case GActorState:
		ga.state = state
		return
	default:
		return
	}
}
