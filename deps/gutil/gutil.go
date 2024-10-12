package gutil

import (
	"fmt"
	"time"
)

// 获取当前时间戳
func Timestamp() int64 {
	return time.Now().Unix()
}

// 格式化当前日期
func FormatDate() string {
	return time.Now().Format("2006-01-02")
}

// 格式化当前时间
func FormatTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 格式化当前时间（毫秒级）
func FormatMilliTime() string {
	// 格式化当前时间，精确到毫秒
	return time.Now().Format("2006/01/02 15:04:05.000")
}

// 格式化当前时间为指定格式
// layout: 2006-01-02_150405_000 | 或者其他
func FormatTimeByLayout(layout string) string {
	return time.Now().Format(layout)
}

// 判断两个时间戳是否是同一天
func IsSameDay(timestamp1, timestamp2 int64) bool {
	t1 := time.Unix(timestamp1, 0)
	t2 := time.Unix(timestamp2, 0)
	return t1.Year() == t2.Year() && t1.YearDay() == t2.YearDay()
}

// 尝试将 interface{} 转换为 string
func ConvToString(data interface{}) string {
	switch v := data.(type) {
	case string:
		return v
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v)
	case float32, float64:
		return fmt.Sprintf("%f", v)
	case bool:
		return fmt.Sprintf("%t", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}
