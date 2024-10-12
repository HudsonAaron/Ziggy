// tcp连接池模块

package gtcp

import (
	"fmt"
	"main/deps/glog"
	"net"
	"sync"
)

// tcp连接池结构体
type TcpPool struct {
	addr    string                // 监听地址
	ts      *TcpServer            // 监听器
	pool    map[net.Conn]*TcpConn // 连接池
	lock    *sync.Mutex           // 连接池锁
	maxSize int                   // 最大连接数
}

var (
	Version = "1.0.0" // 版本号
)

// 获取连接池中的连接
func NewTcpPool(ip string, port int, maxSize int) *TcpPool {
	// 初始化连接池
	return &TcpPool{
		addr:    fmt.Sprintf("%s:%d", ip, port),
		pool:    make(map[net.Conn]*TcpConn),
		lock:    new(sync.Mutex),
		maxSize: maxSize,
	}
}

// 启动连接池
func (tp *TcpPool) Start(msgRouter func(bool, chan []byte)) error {
	if tp.ts == nil {
		// 启动tcp监听
		tp.ts = NewTcpServer(tp.addr)
		err := tp.ts.Start()
		if err != nil {
			glog.Error("TCP Server Start Error:%s", err.Error())
			return err
		}
	}
	go tp.Accept(msgRouter) // 启动接收tcp连接
	glog.Info("TCP Server start up")
	return nil
}

// 启动接收tcp连接
func (tp *TcpPool) Accept(msgRouter func(bool, chan []byte)) {
	for {
		if len(tp.pool) >= tp.maxSize {
			glog.Error("TCP Server Pool Full")
			continue
		}
		conn, err := tp.ts.Accept()
		if err != nil {
			glog.Error("TCP Server Accept Error:%s", err.Error())
			tp.Stop()
			break
		}
		// 将连接放入连接池
		tp.lock.Lock()
		tc := NewTcpConn(conn)
		go tc.RecvMsg()                  // 开始接收消息
		msgRouter(BigEndian, tc.msgChan) // 消息路由
		tp.pool[(*conn)] = tc
		tp.lock.Unlock()
	}
}

// 停止连接池
func (tp *TcpPool) Stop() {
	if tp.ts != nil {
		tp.ts.Stop()
	}
	for _, tc := range tp.pool {
		tc.Stop()
	}
}
