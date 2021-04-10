package func_time

import (
	"business_master"
	"time"
)

func Date() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func SetTimeout(inteval int, fn func(...interface{}) int, args ...interface{}) {
	if inteval < 0 {
		main.LogError("new timerout inteval:%v ms", inteval)
		return
	}
	main.LogInfo("new timerout inteval:%v ms", inteval)

	Go2(func(cstop chan struct{}) {
		var tick *time.Timer
		for inteval > 0 {
			tick = time.NewTimer(time.Millisecond * time.Duration(inteval))
			select {
			case <-cstop:
				tick.Stop()
				inteval = 0
			case <-tick.C:
				tick.Stop()
				inteval = fn(args...)
			}
		}
	})
}
