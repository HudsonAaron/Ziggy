package main

import (
	"fmt"
	"main/deps/glog"
	"main/deps/gtcp"
	"main/src/gconf"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := gconf.LoadConf()
	if err != nil {
		fmt.Println(err)
	}
	// 启动日志
	err = glog.Start(gconf.LogConf)
	if err != nil {
		glog.CrashExit(err.Error())
	}

	// 启动服务
	// ...
	err = gtcp.NewTcpPool(gconf.TcpConf).Start(Init)
	if err != nil {
		glog.CrashExit(err.Error())
	}
	glog.Info("Server start success!")
	// 监听退出信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	msg := <-c
	glog.Crash("Server close. Reason：%v", msg)

	// 关闭服务
	// ...
	// 关闭相关服务完成

	glog.Stop()
}

// 初始化消息路由
func Init(bigEndian bool, msgChan chan []byte) {
	// 业务...
	glog.Info("Init success!")
	for {
		// 处理消息
		<-msgChan
	}
}
