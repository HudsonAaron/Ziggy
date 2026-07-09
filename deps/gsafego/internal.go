package gsafego

import (
	"fmt"
	"main/deps/glog"
	"math/rand"
	"runtime/debug"
	"sync/atomic"
	"time"
)

// idCounter 全局自增计数器，用于生成唯一 goroutine ID
var idCounter int64

// generateID 生成唯一 goroutine ID
// 使用纳秒时间戳 + 原子自增计数器，避免 crypto/rand 的 syscall 开销
func generateID() string {
	return fmt.Sprintf("%d_%d", time.Now().UnixNano(), atomic.AddInt64(&idCounter, 1))
}

// reserveGoroutineSlot 尝试预留一个 goroutine 配额
// 先原子递增再检查上限，超限则回滚。返回 true 表示预留成功
func reserveGoroutineSlot() bool {
	newCount := atomic.AddInt64(&goroutineCount, 1)
	max := atomic.LoadInt64(&gCfg.maxGoroutines)
	if max > 0 && newCount > max {
		// 超限，回滚计数器
		atomic.AddInt64(&goroutineCount, -1)
		return false
	}
	return true
}

// safeRun 是每个安全 goroutine 的执行循环
func safeRun(sg *safeGoroutine) {
	defer close(sg.done)
	defer cleanupGoroutine(sg) // 清理 goroutineMap 和 group

	for {
		// 每次循环前检查是否已被取消
		select {
		case <-sg.ctx.Done():
			glog.Info("[gsafego] goroutine %s 被取消，退出", sg.id)
			return
		default:
		}

		// 在匿名函数中执行用户逻辑，捕获 panic
		panicked := runOnce(sg)

		if !panicked {
			// 正常退出
			return
		}

		// panic 了 → 判断是否可以重启（restartCount 仅本 goroutine 访问，无需锁）
		sg.restartCount++
		restarted := sg.restartCount
		maxRestart := sg.restartMax

		if restarted > maxRestart {
			glog.Error("[gsafego] goroutine %s 已达最大重启次数(%d)，退出",
				sg.id, maxRestart)
			return
		}

		// glog.Info("[gsafego] goroutine %s 第%d次重启（上限%d）",
		// 	sg.id, restarted, maxRestart)

		// 最小重启间隔，防止重启风暴 - 随机化间隔
		delay := time.Duration(rand.Intn(500)) * time.Millisecond
		time.Sleep(delay)
		select {
		case <-time.After(delay):
		case <-sg.ctx.Done():
			return
		}
	}
}

// runOnce 执行一次用户函数，返回是否 panic
func runOnce(sg *safeGoroutine) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
			stack := debug.Stack()
			glog.Crash("[gsafego] goroutine %s panic: %v\nstack:\n%s",
				sg.id, r, string(stack))
		}
	}()

	// 执行用户函数
	sg.fun()
	return false
}

// cleanupGoroutine 从全局管理器移除 goroutine（含释放配额）
func cleanupGoroutine(sg *safeGoroutine) {
	// 释放全局配额
	atomic.AddInt64(&goroutineCount, -1)

	goroutineMap.Delete(sg.id)

	if sg.groupID != "" {
		// 从群组中移除
		if v, ok := groupMap.Load(sg.groupID); ok {
			grp := v.(*safeGroup)
			grp.mu.Lock()
			delete(grp.members, sg.id)
			grp.mu.Unlock()
		}
	}
}
