package role

import (
	"main/deps/gactor"
	"main/deps/glog"
)

// 创建角色
func CreateRole(name string) {
	// TODO: 创建角色
	role, err := gactor.Start(name, nil)
	if err != nil {
		glog.Error("create role failed: %v", err)
	}
	RoleList[name] = role
}

func RoleInfo(v ...interface{}) (interface{}, gactor.GActorState, error) {
	glog.Info("---------%s", v[0])
	return nil, v[len(v)-1], nil
}

// 角色发送消息给另一个角色
func RoleSayHello() {
	RoleList["zhl"].Info(RoleList["hudson"], SayHello)
}

func SayHello(v ...interface{}) (interface{}, gactor.GActorState, error) {
	state := v[len(v)-1]
	glog.Info("hello, i'm " + state.(string))
	friend := v[0].(*gactor.GActor)
	friend.Call(nil, SayHi)
	return nil, state, nil
}

func SayHi(v ...interface{}) (interface{}, gactor.GActorState, error) {
	state := v[len(v)-1]
	glog.Info("hi, i'm " + state.(string))
	return nil, state, nil
}
