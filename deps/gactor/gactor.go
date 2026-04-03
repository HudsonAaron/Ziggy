package gactor

import (
	"fmt"
	"sync"
	"time"
)

var (
	Version      = "1.0.1"                        // 版本号
	MsgQueueSize = 100                            // 消息队列大小
	actorMap     = make(map[string]*gActorStatus) // actor map
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
	actor := &GActor{}
	actorStatus := &gActorStatus{}
	err := doStart(state, handle, actor, actorStatus)
	if err != nil {
		return nil, err
	}
	actorMap[actor.key] = actorStatus
	return actor, nil
}

// 终止actor
func Stop(actorID string, handle Handle) {
	gam := GActorMsg{
		msgType: STOP,
		handle:  handle,
	}
	if actor, ok := actorMap[actorID]; ok && actor != nil && actor.msgChan != nil {
		actor.msgChan <- gam
		actor.msgChan = nil
		actorMap[actorID] = nil
	}
}

// 终止actor，默认超时时间为5秒
func (ga *GActor) Stop(handle Handle) {
	ga.StopTimeout(handle, 5*time.Second)
}

// 终止actor，设置指定超时时间
func (ga *GActor) StopTimeout(handle Handle, timeout time.Duration) {
	mc := make(chan interface{})
	gam := GActorMsg{
		msgType: STOP,
		handle:  handle,
		reply:   mc,
	}
	if actorMap[ga.key] == nil {
		return
	}
	// 发送终止消息到actor的消息队列
	actorStatus := actorMap[ga.key]
	if actorStatus.msgChan != nil {
		actorStatus.msgChan <- gam
	} else {
		close(mc)
		return
	}
	select {
	case <-mc:
		// 关闭回复通道
		close(mc)
		break
	case <-time.After(timeout):
		// 关闭回复通道
		close(mc)
		break
	}
	actorMap[ga.key] = nil
}

// Call函数，默认超时时间为5秒
func (ga *GActor) Call(msg interface{}, handle Handle) (interface{}, error) {
	return ga.CallTimeout(msg, handle, 5*time.Second)
}

// Call函数，设置指定超时时间
func (ga *GActor) CallTimeout(msg interface{}, handle Handle, timeout time.Duration) (interface{}, error) {
	mc := make(chan interface{})
	gam := GActorMsg{
		msgType: CALL,
		msg:     msg,
		handle:  handle,
		reply:   mc,
	}
	// 发送消息到actor的消息队列
	if actorMap[ga.key] == nil {
		close(mc)
		return nil, fmt.Errorf("actor not found")
	}
	actorStatus := actorMap[ga.key]
	if actorStatus.msgChan != nil {
		actorStatus.msgChan <- gam
	} else {
		close(mc)
		return nil, fmt.Errorf("actor not found")
	}
	select {
	case reply := <-mc:
		// 关闭回复通道
		close(mc)
		return reply, nil
	case <-time.After(timeout):
		// 关闭回复通道
		close(mc)
		return nil, fmt.Errorf("actor timeout")
	}
}

// Info函数
func Info(actorID string, msg interface{}, handle Handle) error {
	gam := GActorMsg{
		msgType: INFO,
		msg:     msg,
		handle:  handle,
	}
	if actor, ok := actorMap[actorID]; ok && actor != nil && actor.msgChan != nil {
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
	// 发送消息到actor的消息队列
	if actorMap[ga.key] == nil {
		return fmt.Errorf("actor not found")
	}
	actorStatus := actorMap[ga.key]
	if actorStatus.msgChan != nil {
		actorStatus.msgChan <- gam
	} else {
		return fmt.Errorf("actor not found")
	}
	return nil
}
