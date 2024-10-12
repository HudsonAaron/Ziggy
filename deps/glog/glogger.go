package glog

import (
	"log"
	"main/deps/gutil"
	"os"
	"path/filepath"
)

// 日志结构体
type Logger struct {
	file       *os.File    // 日志文件
	filePath   string      // 日志文件路径
	filename   string      // 日志名称
	path       string      // 日志路径
	level      string      // 日志等级
	maxSize    int64       // 日志文件最大大小
	createTime int64       // 日志创建时间
	msgChan    chan string // 日志消息通道
}

// 日志等级
const (
	INFO    = 1 // info
	WARNING = 2 // warning
	ERROR   = 3 // error
	CRASH   = 4 // crash
)

// 确认日志路径，并创建
func MakeSureDir(filePath string) error {
	// 判断路径是否存在，不存在，则创建路径
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

// 创建指定类型日志文件
func NewLogger(logLv string, logPath string, logName string, maxSize int64) *Logger {
	msgChan := make(chan string) // 创建消息通道
	lg := &Logger{
		level:    logLv,
		filename: logName,
		path:     logPath,
		maxSize:  maxSize,
		msgChan:  msgChan,
	}
	return lg
}

// 打开日志，日志文件不存在，则创建并打开
func (lg *Logger) OpenFile() {
	filePath := GetLogPath(lg.path, lg.filename) // 拼接日志路径
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	lg.file = file
	lg.filePath = filePath
	lg.createTime = gutil.Timestamp()
}

// 启动日志输出
func (lg *Logger) Start() {
	for {
		msg := <-lg.msgChan
		lg.CheckLogStatus()
		lg.file.WriteString(msg + "\n")
	}
}

// 关闭日志文件
func (lg *Logger) Stop() {
	if lg != nil && lg.file != nil {
		defer lg.file.Close()
		lg.file = nil
		lg.createTime = 0
		lg.filePath = ""
	}
}

// 拼接日志路径
func GetLogPath(logPath string, logName string) string {
	dateTime := gutil.FormatTimeByLayout("2006-01-02_150405")
	return filepath.Join(logPath, logName+"_"+dateTime+".log")
}

// 检测日志文件状态
func (lg *Logger) CheckLogStatus() {
	if lg.file != nil {
		fileInfo, err := os.Stat(lg.filePath)
		if err != nil {
			log.Fatal(err)
		}
		// 文件大小超过限制，则关闭文件，重新创建新的日志文件
		IsOverSize := fileInfo.Size() > lg.maxSize
		// 判断当前时间是否跨越天，如果是，则关闭文件，重新创建新的日志文件
		IsSameDay := gutil.IsSameDay(gutil.Timestamp(), lg.createTime)
		if IsOverSize || !IsSameDay {
			lg.Stop() // 关闭文件
		}
	}
	if lg.file == nil { // 文件不存在，则打开文件
		lg.OpenFile()
	}
}
