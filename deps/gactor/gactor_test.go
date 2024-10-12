package gactor

import "main/deps/glog"

func main() {
	actor, err := Start(1, nil)
	if err != nil {
		glog.CrashExit(err.Error())
	}
	reply, err := actor.Call("hello", PrintSS)
	if err != nil {
		glog.CrashExit(err.Error())
	}
	glog.Info(reply.(string))

	// 打印结果
	// ...: sadasd
	// ...: hello
	// ...: world
}

func PrintSS(v ...interface{}) (interface{}, GActorState, error) {
	glog.Info("sadasd")
	glog.Info(v[0].(string))
	// 获取 v 的最后一个元素 为 GActorStater
	return "world", v[len(v)-1], nil
}
