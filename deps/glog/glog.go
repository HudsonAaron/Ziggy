package glog

import (
	"fmt"
	"os"
	"time"
)

var (
	Version   = "1.0.0" // 版本号
	InfoLg    *Logger   // info
	WarningLg *Logger   // warning
	ErrorLg   *Logger   // error
	CrashLg   *Logger   // crash
	LogFormat string    // 日志格式
)

// 打开文件
func Start(logConf map[string]any) error {
	// 设置日志格式
	SetFormat(logConf)
	// 日志文件路径
	filePath, err := GetFilePath(logConf["path"])
	if err != nil {
		return err
	}
	// 日志文件大小限制，默认10M
	logSize, err := GetSize(logConf["size"])
	if err != nil {
		return err
	}
	// 日志等级
	logLv := GetLevel(logConf["level"].(string))
	fileConf, err := GetLogFile(logConf["logfile"])
	if err != nil {
		return err
	}
	for _, lc := range fileConf {
		fLevel := lc.level
		fLv := GetLevel(fLevel)
		fName := lc.filename
		if fLv >= logLv {
			lg := NewLogger(fLevel, filePath, fName, logSize)
			go lg.Start()
			switch lg.level {
			case "info":
				InfoLg = lg
			case "warning":
				WarningLg = lg
			case "error":
				ErrorLg = lg
			case "crash":
				CrashLg = lg
			}
		}
	}
	return nil
}

// 关闭文件
func Stop() {
	time.Sleep(time.Second)
	CrashLg.Stop()
	ErrorLg.Stop()
	WarningLg.Stop()
	InfoLg.Stop()
}

// 写入info日志
func Info(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	module, line := GetModuleLine() // 获取模块名，行号
	logMsg := GetFormat("info", module, line, msg)
	ConsoleLog(logMsg)
	InfoLg.MsgToChan(logMsg)
}

// 写入warning日志
func Warning(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	module, line := GetModuleLine() // 获取模块名，行号
	logMsg := GetFormat("warning", module, line, msg)
	ConsoleLog(logMsg)
	InfoLg.MsgToChan(logMsg)
	WarningLg.MsgToChan(logMsg)
}

// 写入error日志
func Error(format string, v ...any) {
	// 获取调用这个函数的文件名、行号
	msg := fmt.Sprintf(format, v...)
	module, line := GetModuleLine() // 获取模块名，行号
	logMsg := GetFormat("error", module, line, msg)
	ConsoleLog(logMsg)
	InfoLg.MsgToChan(logMsg)
	ErrorLg.MsgToChan(logMsg)
}

// 写入crash日志
func Crash(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	module, line := GetModuleLine() // 获取模块名，行号
	logMsg := GetFormat("crash", module, line, msg)
	ConsoleLog(logMsg)
	InfoLg.MsgToChan(logMsg)
	CrashLg.MsgToChan(logMsg)
}

// 写入crash日志并退出程序
func CrashExit(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	module, line := GetModuleLine() // 获取模块名，行号
	logMsg := GetFormat("crash", module, line, msg)
	ConsoleLog(logMsg)
	InfoLg.MsgToChan(logMsg)
	CrashLg.MsgToChan(logMsg)
	Stop()
	os.Exit(1)
}

// 输出到控制台
func ConsoleLog(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	fmt.Println(msg)
}

// 消息传入消息通道
func (lg *Logger) MsgToChan(msg string) {
	if lg != nil && lg.msgChan != nil {
		lg.msgChan <- msg
	}
}
