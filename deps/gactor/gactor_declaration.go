package gactor

import (
	"sync"
	"time"
)

var (
	Version      = "1.0.4"                        // 版本号
	MsgQueueSize = 100                            // 消息队列大小
	actorMap     = make(map[string]*gActorStatus) // actor map
	lock         sync.RWMutex
)

const (
	INIT        = 0 // init回调类型
	CALL        = 1 // call回调类型
	INFO        = 2 // info回调类型
	STOP        = 3 // stop回调类型
	ADDTIMER    = 4 // add timer回调类型
	CANCELTIMER = 5 // cancel timer回调类型
)

// Actor消息结构体，用于传递消息
type GActorMsg struct {
	reply   chan interface{} // 消息反馈通道
	msgType int              // 消息类型
	msg     interface{}      // 消息内容
	handle  Handle           // 回调函数
	timeout time.Duration    // 超时时间
}

// ActorState结构体
type GActorState interface{}

// ActorTimer结构体
type gActorTimer struct {
	id       string      // timer id
	stopC    chan int    // timer stop channel
	expireAt int64       // timer expire time in milliseconds
	msg      interface{} // timer msg
	handle   Handle      // timer handle
}

// Actor内部结构体，用于存储actor的状态和消息队列
type gActorStatus struct {
	key      string         // actor 主键
	state    GActorState    // actor 状态
	msgChan  chan GActorMsg // actor message queue
	timerMap *sync.Map      // 游离go routine列表 , key: timer id, value: gActorTimer
}

// Actor结构体
type GActor struct {
	key string // actor 主键
}

// 初始化回调函数类型
type HandleInit func(...interface{}) (GActorState, error)

// 通用回调函数类型
type Handle func(...interface{}) (interface{}, GActorState, error)
