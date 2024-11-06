package gtcp

import (
	"main/deps/glog"
)

// 开始监听tcp数据
func (tc *TcpConn) RecvMsg() {
	for {
		buf := make([]byte, 1024)
		n, err := (*tc.conn).Read(buf)
		if err != nil {
			glog.Error("tcp read err:%s", err.Error())
			break
		}
		tc.msg = append(tc.msg, buf[:n]...)
		tm := tc.NewTcpMsg()
		leftMsg, err := tm.Parse()
		if err != nil {
			glog.Error("tcp msg parse err:%s", err.Error())
			break
		}
		tc.msg = leftMsg
		if tm.data != nil {
			// 将解析后的数据放入通道
			tc.msgChan <- tm.data
		}
	}
	tcpPool.Delete(tc) // 从连接池中删除
}
