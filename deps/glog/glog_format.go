package glog

import (
	"fmt"
	"go/importer"
	"main/deps/gutil"
	"runtime"
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

func GetPackageName(importPath string) (string, error) {
	// 使用默认的导入器导入包
	pkg, err := importer.Default().Import(importPath)
	if err != nil {
		return "", err
	}

	// 返回包名
	return pkg.Name(), nil
}

// 获取模块名，行号
func GetModuleLine() (string, int) {
	pc, _, line, _ := runtime.Caller(2) // 2表示上2层
	// file = file[strings.LastIndex(file, "/")+1:]
	name := runtime.FuncForPC(pc).Name()
	return name, line
}
