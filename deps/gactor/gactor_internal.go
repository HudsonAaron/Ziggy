package gactor

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"main/deps/gutil"
	"sync"
)

// 获取actor 主键
func getActorKey(data interface{}) (string, error) {
	// 将数据转换为字节数组
	datastr := gutil.ConvToString(data)
	dataBytes := []byte(datastr)
	// 创建一个新的 MD5 哈希对象
	hash := md5.New()
	// 写入数据到哈希对象
	hash.Write(dataBytes)
	// 计算哈希值
	hashBytes := hash.Sum(nil)
	// 将哈希值转换为十六进制字符串
	actorMD5 := hex.EncodeToString(hashBytes)
	return actorMD5, nil
}

// 初始化结构体
func doStart(state interface{}, handle HandleInit, actor *GActor, actorStatus *gActorStatus) error {
	actorID, err := getActorKey(state)
	if err != nil {
		return err
	}
	_, ok := getActorStatus(actorID)
	if ok {
		return fmt.Errorf("actor already started")
	}
	actor.key = actorID
	actorStatus.state = state
	err = actorStatus.handleInit(handle)
	if err != nil {
		return err
	}
	actorStatus.key = actorID
	actorStatus.msgChan = make(chan GActorMsg, MsgQueueSize)
	actorStatus.timerMap = &sync.Map{}
	go actorStatus.loop()
	return nil
}

// 获取actor状态
func getActorStatus(key string) (*gActorStatus, bool) {
	lock.RLock()
	v, ok := actorMap[key]
	lock.RUnlock()
	return v, ok && v != nil
}

// 设置actor状态
func setActorStatus(key string, v *gActorStatus) {
	lock.Lock()
	actorMap[key] = v
	lock.Unlock()
}

// 删除actor状态
func delActorStatus(key string) {
	lock.Lock()
	delete(actorMap, key)
	lock.Unlock()
}
