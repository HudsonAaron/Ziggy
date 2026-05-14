package glog

import (
	"fmt"
	"os"
	"sync"
)

var (
	Version   = "1.0.3" // 版本号
	infoLg    *Logger   // info
	warningLg *Logger   // warning
	errorLg   *Logger   // error
	crashLg   *Logger   // crash
	logFormat string    // 日志格式
	stopOnce  sync.Once // 确保Stop只执行一次
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
		if GetLevel(lc.level) >= logLv {
			lg := newLogger(lc.level, filePath, lc.filename, logSize)
			lg.Start()
			switch lg.level {
			case "info":
				infoLg = lg
			case "warning":
				warningLg = lg
			case "error":
				errorLg = lg
			case "crash":
				crashLg = lg
			}
		}
	}
	return nil
}

// 关闭文件
func Stop() {
	stopOnce.Do(func() {
		// 安全关闭所有logger（避免nil panic）
		if crashLg != nil {
			crashLg.stop()
		}
		if errorLg != nil {
			errorLg.stop()
		}
		if warningLg != nil {
			warningLg.stop()
		}
		if infoLg != nil {
			infoLg.stop()
		}
	})
}

// 写入info日志
func Info(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	module, line := GetModuleLine() // 获取模块名，行号
	logMsg := GetFormat("info", module, line, msg)
	ConsoleLog(logMsg)
	infoLg.MsgToChan(logMsg)
}

// 写入warning日志
func Warning(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	module, line := GetModuleLine() // 获取模块名，行号
	logMsg := GetFormat("warning", module, line, msg)
	ConsoleLog(logMsg)
	infoLg.MsgToChan(logMsg)
	warningLg.MsgToChan(logMsg)
}

// 写入error日志
func Error(format string, v ...any) {
	// 获取调用这个函数的文件名、行号
	msg := fmt.Sprintf(format, v...)
	module, line := GetModuleLine() // 获取模块名，行号
	logMsg := GetFormat("error", module, line, msg)
	ConsoleLog(logMsg)
	infoLg.MsgToChan(logMsg)
	errorLg.MsgToChan(logMsg)
}

// 写入crash日志
func Crash(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	module, line := GetModuleLine() // 获取模块名，行号
	logMsg := GetFormat("crash", module, line, msg)
	ConsoleLog(logMsg)
	infoLg.MsgToChan(logMsg)
	crashLg.MsgToChan(logMsg)
}

// 写入crash日志并退出程序
func CrashExit(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	module, line := GetModuleLine() // 获取模块名，行号
	logMsg := GetFormat("crash", module, line, msg)
	ConsoleLog(logMsg)
	infoLg.MsgToChan(logMsg)
	crashLg.MsgToChan(logMsg)
	Stop()
	os.Exit(1)
}

// 输出到控制台
func ConsoleLog(format string, v ...any) {
	fmt.Println(format)
}

// 消息传入消息通道
func (lg *Logger) MsgToChan(msg string) {
	if lg != nil && lg.msgChan != nil {
		select {
		case lg.msgChan <- msg:
		default:
			// channel满或已关闭，丢弃避免阻塞
		}
	}
}
