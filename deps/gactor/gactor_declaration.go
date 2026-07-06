package gactor

import (
	"sync"
	"time"
)

// 注：sync.Mutex 用于保护 timers slice 的并发访问。
// startTimerScheduler 和 loop 两个 goroutine 会同时读写 timers，
// 在 GOARCH=386 下 slice header 写入不原子，可能导致撕裂读 → 野指针 → GC 崩溃。

var (
	Version      = "1.1.2" // 版本号
	MsgQueueSize = 100     // 消息队列大小
	actorMap     sync.Map  // actor map
)

const (
	INIT        = 0 // init回调类型
	CALL        = 1 // call回调类型
	INFO        = 2 // info回调类型
	STOP        = 3 // stop回调类型
	ADDTIMER    = 4 // add timer回调类型
	CANCELTIMER = 5 // cancel timer回调类型
)

// ActorState结构体（类型别名，与 any 等价）
type GActorState = any

// replyOnce 用 sync.Once 保护 reply channel，确保最多发送一次
type replyOnce struct {
	ch   chan interface{}
	once sync.Once
}

// Actor消息结构体，用于传递消息
type GActorMsg[S any] struct {
	reply   *replyOnce    // 消息反馈通道（sync.Once 保护）
	msgType int           // 消息类型
	msg     interface{}   // 消息内容
	handle  Handle[S]     // 回调函数
	timeout time.Duration // 超时时间
}

// ActorTimer结构体
type gActorTimer[S any] struct {
	id       string      // timer id
	expireAt int64       // timer expire time in milliseconds
	msg      interface{} // timer msg
	handle   Handle[S]   // timer handle
}

// timerHeap 是最小堆，按 expireAt 排序
type timerHeap[S any] []*gActorTimer[S]

// Actor内部结构体，用于存储actor的状态和消息队列
type gActorStatus[S any] struct {
	key         string            // actor 主键
	state       S                 // actor 状态
	msgChan     chan GActorMsg[S] // actor message queue
	timers      timerHeap[S]      // timer 最小堆
	timerWakeup chan struct{}     // 通知 timer 调度器重新计算等待时间
	timerDone   chan struct{}     // 关闭 timer 调度器
	timerMu     sync.Mutex        // 保护 timers 并发访问（startTimerScheduler 与 loop 竞争）
}

// Actor结构体
type GActor[S any] struct {
	key    string           // actor 主键
	status *gActorStatus[S] // 直接持有状态引用
}

// 初始化回调函数类型
type HandleInit[S any] func(actor *GActor[S], v ...interface{}) (S, error)

// 通用回调函数类型
type Handle[S any] func(actor *GActor[S], v ...interface{}) (interface{}, S, error)
