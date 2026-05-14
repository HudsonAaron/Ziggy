package main

import (
	"fmt"
	"main/deps/ghttp"
	"main/deps/glog"
	"main/src/gconf"
	"main/src/platform"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 加载配置文件
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
	IP, ok := gconf.HttpConf["ip"].(string)
	if !ok {
		glog.CrashExit("HttpConf ip is not string")
	}
	Port, ok := gconf.HttpConf["port"].(float64)
	if !ok {
		glog.CrashExit("HttpConf port is not float64")
	}
	URL := fmt.Sprintf("%s:%.0f", IP, Port)
	ghttp.Start(URL, platform.GetHRouter())

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
