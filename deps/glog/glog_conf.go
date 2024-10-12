package glog

import "fmt"

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

// 获取日志文件配置
func GetLogFile(logFile interface{}) ([]LogFile, error) {
	// 首先，尝试断言 logFile 为 []map[string]string 类型
	var logSlice []map[string]string
	var ok bool
	if logSlice, ok = logFile.([]map[string]string); ok {
		// 如果成功，直接使用 logSlice
	} else {
		// 如果失败，尝试断言为 []interface{} 类型
		var logInterfaceSlice []interface{}
		if logInterfaceSlice, ok = logFile.([]interface{}); !ok {
			return nil, fmt.Errorf("logFile is not a slice")
		}
		// 将 []interface{} 转换为 []map[string]string
		logSlice = make([]map[string]string, len(logInterfaceSlice))
		for i, elem := range logInterfaceSlice {
			lm, ok := elem.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("element at index %d is not a map[string]string", i)
			}
			logMap := make(map[string]string)
			for k, v := range lm {
				logMap[k] = v.(string)
			}
			logSlice[i] = logMap
		}
	}

	// 创建一个结果切片，长度与 logSlice 相同
	result := make([]LogFile, len(logSlice))

	// 遍历 logSlice 中的每个元素
	for i, elem := range logSlice {
		// 从 map 中获取 level 和 filename
		level := elem["level"]
		filename := elem["filename"]

		// 将 LogFile 添加到结果切片中
		result[i] = LogFile{
			level:    level,
			filename: filename,
		}
	}

	return result, nil
}
