package glog

import (
	"fmt"
	"reflect"
)

// 获取日志路径
func GetFilePath(logPath interface{}) (string, error) {
	// 日志文件路径
	filePath := ""
	if logPath != nil {
		filePath = logPath.(string)
		err := MakeSureDir(filePath)
		if err != nil {
			return "", err
		}
	}
	return filePath, nil
}

// 获取日志等级
func GetLevel(level string) int {
	switch level {
	case "info":
		return INFO
	case "warning":
		return WARNING
	case "error":
		return ERROR
	case "crash":
		return CRASH
	default:
		return INFO
	}
}

// 获取日志大小
func GetSize(logSize interface{}) (int64, error) {
	// 日志文件大小限制，默认10M
	size := int64(1024 * 1024 * 10)
	if _size, ok := logSize.(float64); ok {
		size = int64(_size)
	} else if _size, ok := logSize.(int); ok {
		size = int64(_size)
	} else if logSize == nil {
	} else {
		// 处理类型不匹配的情况
		return 0, fmt.Errorf("log size type error")
	}
	return size, nil
}

// 日志文件配置
type LogFile struct {
	level    string
	filename string
}

// 获取日志文件配置（修复类型断言安全问题）
func GetLogFile(logFile interface{}) ([]LogFile, error) {
	if logFile == nil {
		return nil, fmt.Errorf("logFile config is nil")
	}

	// 支持配置文件中 []map[string]any 或 []map[string]string
	v := reflect.ValueOf(logFile)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("logFile is not a slice")
	}

	result := make([]LogFile, v.Len())
	for i := 0; i < v.Len(); i++ {
		item := v.Index(i).Interface()
		m, ok := item.(map[string]interface{})
		if !ok {
			m2, ok2 := item.(map[string]string)
			if !ok2 {
				return nil, fmt.Errorf("element at index %d is not a map", i)
			}
			result[i] = LogFile{
				level:    m2["level"],
				filename: m2["filename"],
			}
			continue
		}
		level, _ := m["level"].(string)
		filename, _ := m["filename"].(string)
		result[i] = LogFile{
			level:    level,
			filename: filename,
		}
	}

	return result, nil
}
