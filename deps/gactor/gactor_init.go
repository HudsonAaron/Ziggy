package gactor

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"main/deps/gutil"
)

// Actor消息结构体，用于传递消息
type GActorMsg struct {
	reply   chan interface{} // 消息反馈通道
	msgType int              // 消息类型
	msg     interface{}      // 消息内容
	handle  Handle           // 回调函数
}

// ActorState结构体
type GActorState interface{}

// Actor结构体
type GActor struct {
	key     string         // actor 主键
	state   GActorState    // actor 状态
	msgChan chan GActorMsg // actor message queue
}

const (
	INIT = 0 // init回调类型
	CALL = 1 // call回调类型
	INFO = 2 // info回调类型
	STOP = 3 // stop回调类型
)

// 获取actor 主键
func GetActorKey(data interface{}) (string, error) {
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
func _Start(state interface{}, handle HandleInit) (*GActor, error) {
	actorID, err := GetActorKey(state)
	if err != nil {
		return nil, err
	}
	if ActorMap[actorID] != nil {
		return ActorMap[actorID], fmt.Errorf("actor already started")
	}
	ga := &GActor{
		key:     actorID,
		msgChan: make(chan GActorMsg, 100),
		state:   state,
	}
	err = ga.HandleInit(handle)
	if err != nil {
		return nil, err
	}
	go ga.Loop()
	return ga, nil
}
