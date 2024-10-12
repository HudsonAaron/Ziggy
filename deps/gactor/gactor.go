package gactor

import (
	"fmt"
	"sync"
)

var (
	Version      = "1.0.0"                  // 版本号
	MsgQueueSize = 100                      // 消息队列大小
	ActorMap     = make(map[string]*GActor) // actor map
	lock         sync.Mutex
)

// 初始化回调函数类型
type HandleInit func(...interface{}) (GActorState, error)

// 通用回调函数类型
type Handle func(...interface{}) (interface{}, GActorState, error)

// 创建actor
func Start(state interface{}, handle HandleInit) (*GActor, error) {
	lock.Lock()
	defer lock.Unlock()
	var actor *GActor
	actor, err := _Start(state, handle)
	if err != nil {
		return actor, err
	}
	ActorMap[actor.key] = actor
	return actor, nil
}

// 终止actor
func Stop(actorID string, handle Handle) {
	gam := GActorMsg{
		msgType: STOP,
		handle:  handle,
	}
	if actor, ok := ActorMap[actorID]; ok && actor != nil && actor.msgChan != nil {
		actor.msgChan <- gam
		actor.msgChan = nil
		ActorMap[actorID] = nil
	}
}

// 终止actor
func (ga *GActor) Stop(handle Handle) {
	gam := GActorMsg{
		msgType: STOP,
		handle:  handle,
	}
	if ga.msgChan != nil {
		ga.msgChan <- gam
		ga.msgChan = nil
	}
	ActorMap[ga.key] = nil
}

// Call函数
func Call(actorID string, msg interface{}, handle Handle) (interface{}, error) {
	mc := make(chan interface{})
	gam := GActorMsg{
		msgType: CALL,
		msg:     msg,
		handle:  handle,
		reply:   mc,
	}
	if actor, ok := ActorMap[actorID]; ok && actor != nil && actor.msgChan != nil {
		actor.msgChan <- gam
	} else {
		close(mc)
		return nil, fmt.Errorf("actor not found")
	}
	reply := <-mc
	return reply, nil
}

// Call函数
func (ga *GActor) Call(msg interface{}, handle Handle) (interface{}, error) {
	mc := make(chan interface{})
	gam := GActorMsg{
		msgType: CALL,
		msg:     msg,
		handle:  handle,
		reply:   mc,
	}
	if ga.msgChan != nil {
		ga.msgChan <- gam
	} else {
		close(mc)
		return nil, fmt.Errorf("actor not found")
	}
	reply := <-mc
	return reply, nil
}

// Info函数
func Info(actorID string, msg interface{}, handle Handle) error {
	gam := GActorMsg{
		msgType: INFO,
		msg:     msg,
		handle:  handle,
	}
	if actor, ok := ActorMap[actorID]; ok && actor != nil && actor.msgChan != nil {
		actor.msgChan <- gam
	} else {
		return fmt.Errorf("actor not found")
	}
	return nil
}

// Info函数
func (ga *GActor) Info(msg interface{}, handle Handle) error {
	gam := GActorMsg{
		msgType: INFO,
		msg:     msg,
		handle:  handle,
	}
	if ga.msgChan != nil {
		ga.msgChan <- gam
	} else {
		return fmt.Errorf("actor not found")
	}
	return nil
}
