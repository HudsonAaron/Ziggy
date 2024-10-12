package glog

import (
	"fmt"
	"main/deps/gutil"
	"runtime"
	"strings"
)

// 拼接日志格式
// "%s【%s:%d】[%s]: %s\n", // time 【module:line】[level]: message
func GetFormat(level string, module string, line int, msg string) string {
	dateTime := gutil.FormatMilliTime()
	return fmt.Sprintf(LogFormat, dateTime, module, line, level, msg)
}

// 设置日志格式
func SetFormat(logConf map[string]any) {
	// 默认格式
	LogFormat = "%s【%s:%d】[%s]: %s\n"
	if logConf["format"] != nil {
		LogFormat = logConf["format"].(string)
	}
}

// 获取模块名，行号
func GetModuleLine() (string, int) {
	_, file, line, _ := runtime.Caller(2) // 2表示上2层
	file = file[strings.LastIndex(file, "/")+1:]
	return file, line
}
