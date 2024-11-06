package gtcp

import (
	"encoding/binary"
	"fmt"
	"main/deps/glog"
	"net"
	"time"
)

func Example() {
	tcpConf := map[string]any{
		"ip":      "127.0.0.1",
		"port":    8080,
		"maxsize": 100,
	}
	err := NewTcpPool(tcpConf).Start(Init)
	if err != nil {
		glog.CrashExit(err.Error())
	}
}

// 初始化消息路由
func Init(bigEndian bool, msgChan chan []byte) {
	// 业务...
}

// / 客户端
func CreateClient() {
	server := "127.0.0.1:9001"
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer conn.Close()
	for {
		time.Sleep(time.Second * 10)
		// hello world 转为大端序二进制
		msg := []byte("hello world")
		// 获取msg长度
		msgLen := len(msg)
		// 将msgLen转为大端序二进制
		msgLenBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(msgLenBytes, uint32(msgLen))
		// 将msgLenBytes和msg拼接
		data := append(msgLenBytes, msg...)
		_, err := conn.Write(data)
		if err != nil {
			fmt.Println("err:", err)
		}
		// fmt.Printf("发送消息：%v\n", data)
	}
}
