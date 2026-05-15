package gactor

import "fmt"

// loop回调
func (ga *gActorStatus[S]) loop() {
	actor := &GActor[S]{key: ga.key, status: ga}
	defer func() {
		close(ga.timerDone)
		delActorStatus(ga.key)
	}()
	for {
		gam := <-ga.msgChan
		switch gam.msgType {
		case ADDTIMER:
			ga.addTimer(&gam)
			continue
		case CANCELTIMER:
			ga.cancelTimer(&gam)
			continue
		case CALL:
			ga.handleCall(actor, &gam)
			continue
		case INFO:
			ga.handleInfo(actor, &gam)
			continue
		case STOP:
			ga.terminate(actor, &gam)
			return
		default:
			fmt.Println("unknown msgType:", gam.msgType)
			return
		}
	}
}
