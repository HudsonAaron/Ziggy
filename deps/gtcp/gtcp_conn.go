package gtcp

import (
	"net"
	"sync"
)

// tcp连接结构体
type TcpConn struct {
	// socket连接
	conn *net.Conn
	// 读写锁
	lock sync.Mutex
	// 消息二进制
	msg []byte
	// 消息通道
	msgChan chan []byte
}

// 初始化tcp连接
func NewTcpConn(conn *net.Conn) *TcpConn {
	return &TcpConn{
		conn:    conn,
		msgChan: make(chan []byte),
	}
}

// 关闭tcp连接
func (tc *TcpConn) Stop() {
	tc.lock.Lock()
	defer tc.lock.Unlock()
	if tc.conn != nil {
		(*tc.conn).Close()
		tc.conn = nil
	}
}

// 获取消息通道
func (tc *TcpConn) GetMsgChan() chan []byte {
	return tc.msgChan
}
