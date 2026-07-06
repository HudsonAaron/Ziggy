package gactor

import (
	"fmt"
	"main/deps/glog"
	"main/deps/gutil"
	"strings"
	"time"
)

// 创建actor
func Start[S any](state S, handle HandleInit[S]) (*GActor[S], error) {
	actor := &GActor[S]{}
	actorStatus := &gActorStatus[S]{}
	err := doStart(state, handle, actor, actorStatus)
	if err != nil {
		return nil, err
	}
	setActorStatus(actor.key, actorStatus)
	return actor, nil
}

// 终止actor（通过 actor 引用，类型安全）
func Stop[S any](actor *GActor[S]) {
	if actor.status == nil {
		return
	}
	mc := make(chan interface{}, 1)
	gam := GActorMsg[S]{
		msgType: STOP,
		reply:   &replyOnce{ch: mc},
	}
	if actor.status.msgChan != nil {
		actor.status.msgChan <- gam
		<-mc
	}
}

// 获取State（类型安全版本）
func GetState[S any](actor *GActor[S]) S {
	if actor.status == nil {
		var zero S
		return zero
	}
	return actor.status.state
}

// 获取status（类型安全版本）
func GetStatus[S any](actor *GActor[S]) string {
	if actor.status == nil {
		return "actor not found"
	}
	actor.status.timerMu.Lock()
	defer actor.status.timerMu.Unlock()
	timerStrs := []string{}
	for _, t := range actor.status.timers {
		timerStrs = append(timerStrs, fmt.Sprintf("{id:%v, expireAt:%v}", t.id, t.expireAt))
	}
	return fmt.Sprintf("{\n\tstate:%v, \n\ttimers:{\n\t\t%v\n}}", actor.status.state, strings.Join(timerStrs, "\n\t\t\t"))
}

// 终止actor，默认超时时间为5秒
func (ga *GActor[S]) Stop(handle Handle[S]) {
	ga.StopTimeout(handle, 5*time.Second)
}

// 终止actor，设置指定超时时间
func (ga *GActor[S]) StopTimeout(handle Handle[S], timeout time.Duration) {
	if ga.status == nil {
		return
	}
	mc := make(chan interface{}, 1)
	gam := GActorMsg[S]{
		msgType: STOP,
		handle:  handle,
		reply:   &replyOnce{ch: mc},
		timeout: timeout,
	}
	if ga.status.msgChan != nil {
		ga.status.msgChan <- gam
	}
	select {
	case <-mc:
	case <-time.After(timeout):
	}
}

// Call函数，默认超时时间为5秒
func (ga *GActor[S]) Call(msg interface{}, handle Handle[S]) (interface{}, error) {
	return ga.CallTimeout(msg, handle, 5*time.Second)
}

// Call函数，设置指定超时时间
func (ga *GActor[S]) CallTimeout(msg interface{}, handle Handle[S], timeout time.Duration) (interface{}, error) {
	if ga.status == nil {
		return nil, fmt.Errorf("actor not found")
	}
	mc := make(chan interface{}, 1)
	gam := GActorMsg[S]{
		msgType: CALL,
		msg:     msg,
		handle:  handle,
		reply:   &replyOnce{ch: mc},
	}
	if ga.status.msgChan != nil {
		ga.status.msgChan <- gam
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
func (ga *GActor[S]) Info(msg interface{}, handle Handle[S]) error {
	if ga.status == nil {
		return fmt.Errorf("actor not found")
	}
	gam := GActorMsg[S]{
		msgType: INFO,
		msg:     msg,
		handle:  handle,
	}
	if ga.status.msgChan != nil {
		ga.status.msgChan <- gam
	} else {
		return fmt.Errorf("actor not found")
	}
	return nil
}

// 创建timer, 内部走Info函数
func (ga *GActor[S]) SendAfter(timeout time.Duration, msg interface{}, handle Handle[S]) (string, error) {
	if ga.status == nil {
		return "", fmt.Errorf("actor not found")
	}
	gTimer := &gActorTimer[S]{
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
	gam := GActorMsg[S]{
		msgType: ADDTIMER,
		timeout: timeout,
		msg:     gTimer,
	}
	if ga.status.msgChan != nil {
		ga.status.msgChan <- gam
	} else {
		return "", fmt.Errorf("actor not found")
	}
	return timerId, nil
}

// 取消timer
func (ga *GActor[S]) CancelTimer(timerID string) {
	if ga.status == nil {
		return
	}
	gam := GActorMsg[S]{
		msgType: CANCELTIMER,
		msg:     timerID,
	}
	if ga.status.msgChan != nil {
		ga.status.msgChan <- gam
	}
}

// 获取timer剩余时间
func (ga *GActor[S]) GetTimerTTL(timerID string) time.Duration {
	if ga.status == nil {
		return 0
	}
	ga.status.timerMu.Lock()
	defer ga.status.timerMu.Unlock()
	for _, t := range ga.status.timers {
		if t.id == timerID {
			return time.Duration(t.expireAt-gutil.TimestampMilli()) * time.Millisecond
		}
	}
	return 0
}
