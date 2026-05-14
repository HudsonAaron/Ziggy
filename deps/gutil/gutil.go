package gutil

import (
	"fmt"
	"math/rand"
	"time"
)

// 获取当前时间戳-秒
func Timestamp() int64 {
	return time.Now().Unix()
}

// 获取当前时间戳-毫秒
func TimestampMilli() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// 格式化当前日期
func FormatDate(v ...interface{}) string {
	if len(v) == 0 {
		return time.Now().Format("2006-01-02")
	}
	// 格式化指定时间，精确到秒
	return time.Unix(v[0].(int64), 0).Format("2006-01-02")
}

// 格式化当前时间
func FormatTime(v ...interface{}) string {
	if len(v) == 0 {
		return time.Now().Format("2006-01-02 15:04:05")
	}
	// 格式化指定时间，精确到秒
	return time.Unix(v[0].(int64), 0).Format("2006-01-02 15:04:05")
}

// 格式化当前时间（毫秒级）
func FormatMilliTime(v ...interface{}) string {
	// 格式化当前时间，精确到毫秒
	if len(v) == 0 {
		return time.Now().Format("2006/01/02 15:04:05.000")
	}
	// 格式化指定时间，精确到毫秒
	return time.Unix(v[0].(int64), 0).Format("2006/01/02 15:04:05.000")
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

// 判断元素是否包含在列表中
func Contains[T comparable](items []T, item T) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}

// 将列表打乱顺序
func ShuffleList[T any](items []T) []T {
	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})
	return items
}

// 获取从min到max之间的随机数
func GetRandomInt(min, max int) int {
	if min >= max {
		return min
	}
	return min + rand.Intn(max-min+1)
}

// 列表中随机抽选1个元素
func GetRandomItem[T any](items []T) T {
	if len(items) == 0 {
		var zero T
		return zero
	}
	return items[rand.Intn(len(items))]
}

// 列表中随机抽选n个元素，可重复
func GetRepeatRandomItems[T any](items []T, n int) []T {
	if len(items) == 0 {
		return []T{}
	}
	result := make([]T, n)
	for i := 0; i < n; i++ {
		result[i] = items[rand.Intn(len(items))]
	}
	return result
}

// 列表中随机抽选n个元素，不可重复
func GetRandomItems[T any](items []T, n int) []T {
	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})
	if n >= len(items) {
		return items
	}
	return items[:n]
}

// 列表中根据权重随机抽选1个元素
// weightFunc: 用于获取每个元素的权重的函数
func GetRandomItemByWeight[T any](items []T, weightFunc func(T) int) T {
	if len(items) == 0 {
		var zero T
		return zero
	}

	// 计算总权重
	totalWeight := 0
	for _, item := range items {
		totalWeight += weightFunc(item)
	}

	if totalWeight <= 0 {
		// 如果总权重为0，随机返回一个
		return items[rand.Intn(len(items))]
	}

	// 生成随机数
	random := rand.Intn(totalWeight)

	// 根据权重选择元素
	currentWeight := 0
	for _, item := range items {
		currentWeight += weightFunc(item)
		if random < currentWeight {
			return item
		}
	}

	// 理论上不会到达这里，保险起见返回第一个元素
	return items[0]
}

// 从列表中移除指定元素
func RemoveItem[T comparable](items []T, item T) []T {
	for i := 0; i < len(items); i++ {
		if items[i] == item {
			items = append(items[:i], items[i+1:]...)
			i--
			break
		}
	}
	return items
}

// 从列表中移除所有指定元素
func RemoveAllItems[T comparable](items []T, item T) []T {
	for i := 0; i < len(items); i++ {
		if items[i] == item {
			items = append(items[:i], items[i+1:]...)
			i--
			break
		}
	}
	return items
}

// 从列表中移除所有指定元素
func RemoveAllItemsByValue[T comparable](items []T, value T) []T {
	for i := 0; i < len(items); i++ {
		if items[i] == value {
			items = append(items[:i], items[i+1:]...)
			i--
			break
		}
	}
	return items
}
