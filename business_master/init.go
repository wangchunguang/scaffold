package main

import (
	"runtime"
)

var DefLog *Log //日志

func init() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	DefLog = NewLog(10000, &ConsoleLogger{true})
	DefLog.SetLevel(LogLevelInfo)

}
