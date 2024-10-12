// tcp服务
package gtcp

import (
	"net"
)

// 监听器结构体
type TcpServer struct {
	// 服务器地址
	addr string
	// 服务器监听器
	listener *net.Listener
}

// 初始化tcp服务器
func NewTcpServer(addr string) *TcpServer {
	return &TcpServer{
		addr:     addr,
		listener: nil,
	}
}

// 创建tcp服务器
func (ts *TcpServer) Start() error {
	var err error
	// 启动Tcp监听
	listener, err := net.Listen("tcp", ts.addr)
	if err != nil {
		return err
	}
	ts.listener = &listener
	return nil
}

// 接受tcp连接
func (ts *TcpServer) Accept() (*net.Conn, error) {
	conn, err := (*ts.listener).Accept()
	if err != nil {
		return nil, err
	}
	return &conn, nil
}

// 关闭tcp服务器
func (ts *TcpServer) Stop() error {
	return (*ts.listener).Close()
}
