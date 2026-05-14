package gactor

import "fmt"

// loop回调
func (ga *gActorStatus) loop() {
	// 退出时清理timer和actor状态
	defer func() {
		ga.timerMap.Range(func(key, value interface{}) bool {
			timer := value.(*gActorTimer)
			if timer.stopC != nil {
				close(timer.stopC)
			}
			ga.timerMap.Delete(key)
			return true
		})
		delActorStatus(ga.key)
	}()
	// 循环处理消息
	for {
		gam := <-ga.msgChan
		// 处理消息
		switch gam.msgType {
		case ADDTIMER: // add timer回调类型
			ga.addTimer(&gam)
			continue
		case CANCELTIMER: // cancel timer回调类型
			ga.cancelTimer(&gam)
			continue
		case CALL: // call回调类型
			ga.handleCall(&gam)
			continue
		case INFO: // info回调类型
			ga.handleInfo(&gam)
			continue
		case STOP: // stop回调类型
			ga.terminate(&gam)
			return
		default:
			fmt.Println("unknown msgType:", gam.msgType)
			return
		}
	}
}
