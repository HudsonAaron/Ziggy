// tcp连接池模块

package gtcp

import (
	"fmt"
	"main/deps/glog"
	"net"
	"sync"
	"time"
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
	Version = "1.0.0"    // 版本号
	tcpPool = &TcpPool{} // tcp连接池
)

// 获取连接池中的连接
func NewTcpPool(tcpConf map[string]any) *TcpPool {
	ip, port, maxSize, err := GetTcpConf(tcpConf)
	if err != nil {
		glog.CrashExit("TCP Server GetConf Error:%s", err.Error())
		return nil
	}
	if maxSize <= 0 {
		glog.CrashExit("TCP Server MaxSize too small~")
		return nil
	}
	// 初始化连接池
	tcpPool = &TcpPool{
		addr:    fmt.Sprintf("%s:%d", ip, port),
		pool:    make(map[net.Conn]*TcpConn),
		lock:    new(sync.Mutex),
		maxSize: maxSize,
	}
	return tcpPool
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
	go tp.ShowConnCount()   // 启动定时统计连接数量
	glog.Info("TCP Server start up")
	return nil
}

// 启动接收tcp连接
func (tp *TcpPool) Accept(msgRouter func(bool, chan []byte)) {
	for {
		if len(tp.pool) >= tp.maxSize {
			glog.Error("TCP Server Pool Full")
			time.Sleep(time.Second * 10) // 连接池满时，等待10秒再接收新的连接
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
		tp.pool[(*conn)] = tc
		tp.lock.Unlock()
		go tc.RecvMsg()                     // 开始接收消息
		go msgRouter(BigEndian, tc.msgChan) // 消息路由
	}
}

// 停止连接池
func (tp *TcpPool) Stop() {
	if tp.ts != nil {
		tp.ts.Stop()
	}
	for _, tc := range tp.pool {
		tp.lock.Lock()
		delete(tp.pool, (*tc.conn))
		tc.Stop()
		tp.lock.Unlock()
	}
}

// 关闭连接池中的连接
func (tp *TcpPool) Delete(tc *TcpConn) {
	tp.lock.Lock()
	tc, ok := tp.pool[(*tc.conn)]
	if ok {
		delete(tp.pool, (*tc.conn))
		tc.Stop()
	}
	tp.lock.Unlock()
}

// 获取连接池配置
func GetTcpConf(tcpConf map[string]any) (string, int, int, error) {
	var ip string = "127.0.0.1" // 默认监听地址
	var port int
	var maxSize int = 100 // 默认最大连接数100
	if tcpConf["ip"] != nil {
		if _ip, ok := tcpConf["ip"].(string); ok {
			ip = _ip
		} else {
			return "", 0, 0, fmt.Errorf("tcpConf ip type error")
		}
	}
	if tcpConf["port"] != nil {
		if _port, ok := tcpConf["port"].(int); ok {
			port = _port
		} else if _port, ok := tcpConf["port"].(float64); ok {
			port = int(_port)
		} else {
			return "", 0, 0, fmt.Errorf("tcpConf port type error")
		}
	} else {
		return "", 0, 0, fmt.Errorf("tcpConf port not found")
	}
	if tcpConf["maxsize"] != nil {
		if _maxSize, ok := tcpConf["maxsize"].(int); ok {
			maxSize = _maxSize
		} else if _maxSize, ok := tcpConf["maxsize"].(float64); ok {
			maxSize = int(_maxSize)
		} else {
			return "", 0, 0, fmt.Errorf("tcpConf maxSize type error")
		}
	}
	return ip, port, maxSize, nil
}

// 定时统计连接数量
func (tp *TcpPool) ShowConnCount() {
	for {
		glog.Info("TCP Server current conn count:%d maxSize:%d", len(tp.pool), tp.maxSize)
		<-time.After(10 * time.Second)
	}
}
