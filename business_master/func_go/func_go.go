package func_go

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var stop int32 //停止标志
var stopChanForLog = make(chan struct{})
var syncGroup sync.WaitGroup
var poolChan = make(chan func())
var poolGoCount int32

func Go(fn func()) {
	if poolGoCount > 0 {
		Go2(fn)
	}
	select {
	case poolChan <- fn:
		return
	default:
		atomic.AddInt32(&poolGoCount, 1)
	}
}

func Go2(fn func()) {
	atomic.AddInt32(&poolGoCount, -1)
	syncGroup.Add(1)
	go func() {

	}()

}

func Try(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error catch", err)
		}
	}()
	fn()
}

func GoForLog(fn func(cstop chan struct{})) bool {
	if IsStop() {
		return false
	}
	syncGroup.Add(1)

	go func() {
		fn(stopChanForLog)
		syncGroup.Done()
	}()
	return true
}
func IsStop() bool {
	return stop == 1
}
