package logs

import (
	rotatelogs "github.com/lestrrat-func_go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

var stdFormatter *prefixed.TextFormatter  // 命令行输出格式
var fileFormatter *prefixed.TextFormatter // 文件输出格式

func init() {
	stdFormatter = &prefixed.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02.15:04:05",
		ForceFormatting: true,
		ForceColors:     true,
		DisableColors:   false,
	}
	fileFormatter = &prefixed.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02.15:04:05",
		ForceFormatting: true,
		ForceColors:     false,
		DisableColors:   true,
	}
	//设置标准的记录器格式化程序。
	log.SetFormatter(stdFormatter)
	// 设置日志级别
	log.SetLevel(log.DebugLevel)
}

func InitLog(path string) {

	// 获取当前文件的位置
	file, _ := os.Getwd()
	//日志输出文件位置
	writer, _ := rotatelogs.New(file + path + "%Y%m%d")
	hook := lfshook.NewHook(lfshook.WriterMap{
		log.InfoLevel:  writer,
		log.DebugLevel: writer,
		log.ErrorLevel: writer,
	}, fileFormatter)
	log.SetOutput(os.Stdout)
	log.AddHook(hook)
	log.Info("something ....")
}
