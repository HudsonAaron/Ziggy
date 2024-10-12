package gproto

import (
	"fmt"
	"main/deps/glog"
)

// 消息路由接口
type GInternal interface {
	GRouter
	Init()
}

// 消息路由
type GRouter struct {
	// 消息类型 大端序 | 小端序
	bigEndian bool
	// 消息通道
	msgChan chan []byte
}

// 初始化消息路由
func Init(bigEndian bool, msgChan chan []byte) {
	gr := GRouter{
		bigEndian: bigEndian,
		msgChan:   msgChan,
	}
	go gr.MsgRouter()
}

// 消息路由
func (gr *GRouter) MsgRouter() {
	for {
		msgBytes := <-gr.msgChan
		if len(msgBytes) == 0 {
			glog.Error("MsgRouter msgBytes is nil")
			continue
		}
		err := ParseProto(msgBytes)
		if err != nil {
			glog.Error("MsgRouter ParseProto err:%s", err.Error())
			continue
		}
	}
}

// 协议路由
func ProtoRouter(protoID int32, data []byte) {
	switch protoID {
	case 1:
		fmt.Println("ProtoRouter:", string(data))
	default:
		fmt.Println("ProtoRouter protoID:", protoID, "data:", string(data))
	}
}
