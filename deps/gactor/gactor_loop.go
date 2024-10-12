package gactor

import "fmt"

// loop回调
func (ga *GActor) Loop() {
	for {
		gam := <-ga.msgChan
		fmt.Println("Loop:", ga.state)
		// 处理消息
		switch gam.msgType {
		case CALL: // call回调类型
			ga.HandleCall(&gam)
			continue
		case INFO: // info回调类型
			ga.HandleInfo(&gam)
			continue
		case STOP: // stop回调类型
			ga.Terminate(&gam)
			return
		}
	}
}
