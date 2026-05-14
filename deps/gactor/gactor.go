package gactor

import (
	"fmt"
	"main/deps/glog"
	"main/deps/gutil"
	"strings"
	"time"
)

// 创建actor
func Start(state interface{}, handle HandleInit) (*GActor, error) {
	actor := &GActor{}
	actorStatus := &gActorStatus{}
	err := doStart(state, handle, actor, actorStatus)
	if err != nil {
		return nil, err
		return nil, err
	}
	setActorStatus(actor.key, actorStatus)
	return actor, nil
}

// 终止actor
func Stop(actorID string, handle Handle) {
	gam := GActorMsg{
		msgType: STOP,
		handle:  handle,
	}
	actorStatus, ok := getActorStatus(actorID)
	if !ok {
		return
	}
	if actorStatus.msgChan != nil {
		actorStatus.msgChan <- gam
	}
}

// 获取State
func GetState(actor *GActor) interface{} {
	actorStatus, ok := getActorStatus(actor.key)
	if !ok {
		return nil
	}
	return actorStatus.state
}

// 获取status
func GetStatus(actor *GActor) string {
	actorStatus, ok := getActorStatus(actor.key)
	if !ok {
		return "actor not found"
	}
	timerMapStr := []string{}
	actorStatus.timerMap.Range(func(key, value interface{}) bool {
		timerMapStr = append(timerMapStr, fmt.Sprintf("{key:%v, value:%v}", key, value))
		return true
	})
	return fmt.Sprintf("{\n\tstate:%v, \n\ttimerMap:{\n\t\t%v\n}}", actorStatus.state, strings.Join(timerMapStr, "\n\t\t\t"))
}

// 终止actor，默认超时时间为5秒
func (ga *GActor) Stop(handle Handle) {
	ga.StopTimeout(handle, 5*time.Second)
}

// 终止actor，设置指定超时时间
func (ga *GActor) StopTimeout(handle Handle, timeout time.Duration) {
	actorStatus, ok := getActorStatus(ga.key)
	if !ok {
		return
	}
	mc := make(chan interface{}, 1)
	gam := GActorMsg{
		msgType: STOP,
		handle:  handle,
		reply:   mc,
		timeout: timeout,
	}
	if actorStatus.msgChan != nil {
		actorStatus.msgChan <- gam
	} else {
		return
	}
	select {
	case <-mc:
		break
	case <-time.After(timeout):
		break
	}
}

// Call函数，默认超时时间为5秒
// Call函数，默认超时时间为5秒
func (ga *GActor) Call(msg interface{}, handle Handle) (interface{}, error) {
	return ga.CallTimeout(msg, handle, 5*time.Second)
}

// Call函数，设置指定超时时间
func (ga *GActor) CallTimeout(msg interface{}, handle Handle, timeout time.Duration) (interface{}, error) {
	// 发送消息到actor的消息队列
	actorStatus, ok := getActorStatus(ga.key)
	if !ok {
		return nil, fmt.Errorf("actor not found")
	}
	mc := make(chan interface{}, 1)
	gam := GActorMsg{
		msgType: CALL,
		msg:     msg,
		handle:  handle,
		reply:   mc,
	}
	if actorStatus.msgChan != nil {
		actorStatus.msgChan <- gam
	} else {
		return nil, fmt.Errorf("actor not found")
	}
	select {
	case reply := <-mc:
		return reply, nil
	case <-time.After(timeout):
		return nil, fmt.Errorf("actor timeout")
	}
}

// Info函数
func (ga *GActor) Info(msg interface{}, handle Handle) error {
	// 发送消息到actor的消息队列
	actorStatus, ok := getActorStatus(ga.key)
	if !ok {
		return fmt.Errorf("actor not found")
	}
	gam := GActorMsg{
		msgType: INFO,
		msg:     msg,
		handle:  handle,
	}
	if actorStatus.msgChan != nil {
		actorStatus.msgChan <- gam
	} else {
		return fmt.Errorf("actor not found")
	}
	return nil
}

// 创建timer, 内部走Info函数
func (ga *GActor) SendAfter(timeout time.Duration, msg interface{}, handle Handle) (string, error) {
	// 发送消息到actor的消息队列
	actorStatus, ok := getActorStatus(ga.key)
	if !ok {
		return "", fmt.Errorf("actor not found")
	}
	// 创建timer
	gTimer := &gActorTimer{
		id:       fmt.Sprintf("%v", msg),
		stopC:    make(chan int),
		expireAt: gutil.TimestampMilli() + timeout.Milliseconds(),
		msg:      msg,
		handle:   handle,
	}
	timerId, err := getActorKey(gTimer)
	if err != nil {
		glog.Error("addTimer err: %v", err)
		return "", err
	}
	gTimer.id = timerId
	// 添加timer
	gam := GActorMsg{
		msgType: ADDTIMER,
		timeout: timeout,
		msg:     gTimer,
	}
	if actorStatus.msgChan != nil {
		actorStatus.msgChan <- gam
	} else {
		return "", fmt.Errorf("actor not found")
	}
	return timerId, nil
}

// 取消timer
func (ga *GActor) CancelTimer(timerID string) {
	// 发送消息到actor的消息队列
	actorStatus, ok := getActorStatus(ga.key)
	if !ok {
		return
	}
	gam := GActorMsg{
		msgType: CANCELTIMER,
		msg:     timerID,
	}
	if actorStatus.msgChan != nil {
		actorStatus.msgChan <- gam
	} else {
		return
	}
}

// 获取timer剩余时间
func (ga *GActor) GetTimerTTL(timerID string) time.Duration {
	actorStatus, ok := getActorStatus(ga.key)
	if !ok {
		return 0
	}
	timer, ok := actorStatus.timerMap.Load(timerID)
	if !ok {
		return 0
	}
	return time.Duration(timer.(*gActorTimer).expireAt-gutil.TimestampMilli()) * time.Millisecond
}
