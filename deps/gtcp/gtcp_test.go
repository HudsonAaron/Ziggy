package gtcp

import (
	"main/deps/glog"
)

func Example() {
	err := NewTcpPool("127.0.0.1", 10001, 1000).Start(Init)
	if err != nil {
		glog.CrashExit(err.Error())
	}
}

// 初始化消息路由
func Init(bigEndian bool, msgChan chan []byte) {
	// 业务...
}
