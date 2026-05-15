package gactor

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"main/deps/gutil"
)

// 获取actor 主键
func getActorKey(data interface{}) (string, error) {
	datastr := gutil.ConvToString(data)
	dataBytes := []byte(datastr)
	hash := md5.New()
	hash.Write(dataBytes)
	hashBytes := hash.Sum(nil)
	actorMD5 := hex.EncodeToString(hashBytes)
	return actorMD5, nil
}

// 初始化结构体
func doStart[S any](state S, handle HandleInit[S], actor *GActor[S], actorStatus *gActorStatus[S]) error {
	actorID, err := getActorKey(state)
	if err != nil {
		return err
	}
	_, ok := getAnyActorStatus(actorID)
	if ok {
		return fmt.Errorf("actor already started")
	}
	actor.key = actorID
	actor.status = actorStatus
	actorStatus.key = actorID
	actorStatus.state = state
	actorStatus.msgChan = make(chan GActorMsg[S], MsgQueueSize)
	actorStatus.timers = make(timerHeap[S], 0)
	actorStatus.timerWakeup = make(chan struct{}, 1)
	actorStatus.timerDone = make(chan struct{})
	err = actorStatus.handleInit(actor, handle)
	if err != nil {
		return err
	}
	go actorStatus.startTimerScheduler()
	go actorStatus.loop()
	return nil
}

func (r *replyOnce) send(data interface{}) {
	r.once.Do(func() {
		select {
		case r.ch <- data:
		default:
			// channel 满了或已关闭，安全丢弃
		}
	})
}

// timerHeap定时器数量
func (h timerHeap[S]) Len() int { return len(h) }

// timerHeap定时器比较
func (h timerHeap[S]) Less(i, j int) bool { return h[i].expireAt < h[j].expireAt }

// timerHeap定时器交换
func (h timerHeap[S]) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// timerHeap定时器添加
func (h *timerHeap[S]) Push(x any) { *h = append(*h, x.(*gActorTimer[S])) }

// timerHeap定时器弹出
func (h *timerHeap[S]) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// 获取actor状态（类型安全版本）
func getActorStatus[S any](key string) (*gActorStatus[S], bool) {
	v, ok := actorMap.Load(key)
	if !ok {
		return nil, false
	}
	return v.(*gActorStatus[S]), true
}

// 获取actor状态（任意类型版本，用于 key 查找）
func getAnyActorStatus(key string) (any, bool) {
	return actorMap.Load(key)
}

// 设置actor状态
func setActorStatus[S any](key string, v *gActorStatus[S]) {
	actorMap.Store(key, v)
}

// 删除actor状态
func delActorStatus(key string) {
	actorMap.Delete(key)
}
