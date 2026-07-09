package gsafego

import (
	"context"
	"sync"
)

// ────────── 常量 ──────────

const (
	Version              = "1.0.0"
	DefaultRestartMax    = 5      // 默认最大重启次数
	DefaultMaxGoroutines = 100000 // 默认全局最大 goroutine 数量（10w）
)

// ────────── 全局配置 ──────────

// gConfig 全局配置，由 Init 设置
type gConfig struct {
	maxGoroutines int64 // 最大 goroutine 数量上限（原子操作，0 = 未限制）
}

var (
	gCfg = gConfig{
		maxGoroutines: DefaultMaxGoroutines,
	} // 全局配置实例
	goroutineCount int64 // 当前活跃 goroutine 数量（原子操作）
)

// ────────── 单 goroutine ──────────

// safeGoroutine 表示一个安全 goroutine 的运行时状态
type safeGoroutine struct {
	id           string             // 唯一标识
	groupID      string             // 所属群组 ID（空字符串 = 独立模式）
	restartCount int                // 当前已重启次数（仅 safeRun goroutine 访问）
	restartMax   int                // 最大允许重启次数（首次执行不计入重启）
	fun          func()             // 用户要执行的函数
	ctx          context.Context    // 取消上下文
	cancel       context.CancelFunc // 取消函数
	done         chan struct{}      // goroutine 完成信号（关闭时发送）
}

// ────────── 群组 ──────────

// safeGroup 表示一个 goroutine 群组
type safeGroup struct {
	id      string                    // 群组 ID
	members map[string]*safeGoroutine // 群组成员（key = goroutineID）
	mu      sync.RWMutex              // 保护 members 并发读写
}

// ────────── 全局管理器 ──────────

var (
	goroutineMap sync.Map // map[string]*safeGoroutine   (key = goroutineID)
	groupMap     sync.Map // map[string]*safeGroup       (key = groupID)
)
