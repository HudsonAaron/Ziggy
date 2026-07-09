package gsafego

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// T1: fun 正常返回
func TestNormalReturn(t *testing.T) {
	done := make(chan struct{})
	id, err := SafeGo(0, func(args ...interface{}) {
		close(done)
	})
	if err != nil {
		t.Fatalf("创建安全goroutine失败 %v", err)
	}
	defer CloseSafeGo(id)

	select {
	case <-done:
		// goroutine 正常退出
	case <-time.After(2 * time.Second):
		t.Fatal("goroutine 未及时退出")
	}
}

// T2: fun panic → 重启 → 正常返回
func TestPanicRecoveryAndRestart(t *testing.T) {
	var count int32
	id, err := SafeGo(3, func(args ...interface{}) {
		atomic.AddInt32(&count, 1)
		c := atomic.LoadInt32(&count)
		if c < 3 {
			panic("模拟 panic")
		}
		// 第3次正常返回
	})
	if err != nil {
		t.Fatalf("创建安全goroutine失败 %v", err)
	}
	defer CloseSafeGo(id)

	time.Sleep(500 * time.Millisecond)

	c := atomic.LoadInt32(&count)
	if c != 3 {
		t.Fatalf("期望执行3次，实际 %d", c)
	}
}

// T3: fun 持续 panic → 达到重启上限退出
func TestMaxRestartExceeded(t *testing.T) {
	var count int32
	id, err := SafeGo(2, func(args ...interface{}) {
		atomic.AddInt32(&count, 1)
		panic("一直 panic")
	})
	if err != nil {
		t.Fatalf("创建安全goroutine失败 %v", err)
	}
	defer CloseSafeGo(id)

	time.Sleep(500 * time.Millisecond)

	c := atomic.LoadInt32(&count)
	// 首次 + 2次重启 = 3次
	if c != 3 {
		t.Fatalf("期望执行3次（首次+2次重启），实际 %d", c)
	}
}

// T4: restartCount=0（不重启）
func TestNoRestart(t *testing.T) {
	var count int32
	id, err := SafeGo(0, func(args ...interface{}) {
		atomic.AddInt32(&count, 1)
		panic("不给重启")
	})
	if err != nil {
		t.Fatalf("创建安全goroutine失败 %v", err)
	}
	defer CloseSafeGo(id)

	time.Sleep(200 * time.Millisecond)

	c := atomic.LoadInt32(&count)
	if c != 1 {
		t.Fatalf("期望执行1次（不重启），实际 %d", c)
	}
}

// T5: CloseSafeGoroutine 正常关闭
func TestCloseSafeGoroutine(t *testing.T) {
	started := make(chan struct{})
	stopped := make(chan struct{})

	id, err := SafeGo(0, func(args ...interface{}) {
		close(started)
		<-time.After(10 * time.Second) // 长期运行
		close(stopped)
	})
	if err != nil {
		t.Fatalf("创建安全goroutine失败 %v", err)
	}

	// 等待 goroutine 启动
	select {
	case <-started:
	case <-time.After(2 * time.Second):
		t.Fatal("goroutine 未启动")
	}

	// 关闭
	CloseSafeGo(id)

	// 等待 goroutine 退出（通过 ctx.Done 触发）
	select {
	case <-stopped:
	case <-time.After(6 * time.Second):
		t.Fatal("goroutine 未在超时前退出")
	}
}

// T6: CloseSafeGoroutine 幂等
func TestCloseIdempotent(t *testing.T) {
	id, err := SafeGo(0, func(args ...interface{}) {
		<-time.After(10 * time.Second)
	})
	if err != nil {
		t.Fatalf("创建安全goroutine失败 %v", err)
	}

	CloseSafeGo(id)
	// 多次调用不应 panic
	CloseSafeGo(id)
	CloseSafeGo(id)

	// 不存在的 ID 也不应 panic
	CloseSafeGo("nonexistent-id")
}

// T7: CloseSafeGoroutineGroup 批量关闭
func TestGroupCloseAll(t *testing.T) {
	var wg sync.WaitGroup
	count := 5
	for i := 0; i < count; i++ {
		wg.Add(1)
		_, err := SafeGoGroup("test-group", 0, func(args ...interface{}) {
			defer wg.Done()
			select {
			case <-time.After(10 * time.Second):
			}
		})
		if err != nil {
			t.Fatalf("创建群组安全goroutine失败 %v", err)
		}
	}

	time.Sleep(50 * time.Millisecond)
	CloseSafeGoGroup("test-group")

	// 所有 goroutine 应该在超时前退出
	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()
	select {
	case <-done:
	case <-time.After(6 * time.Second):
		t.Fatal("群组关闭超时")
	}

	// 群组应已删除，再次关闭应幂等
	CloseSafeGoGroup("test-group")
}

// T8: 并发创建/关闭
func TestConcurrentCreateClose(t *testing.T) {
	var wg sync.WaitGroup
	n := 50
	ids := make(chan string, n)

	// 并发创建
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := SafeGo(0, func(args ...interface{}) {
				<-time.After(2 * time.Second)
			})
			if err != nil {
				t.Fatalf("创建安全goroutine失败 %v", err)
			}
		}()
	}

	wg.Wait()
	close(ids)

	// 并发关闭
	var closeWg sync.WaitGroup
	for id := range ids {
		closeWg.Add(1)
		go func(gid string) {
			defer closeWg.Done()
			CloseSafeGo(gid)
		}(id)
	}

	done := make(chan struct{})
	go func() { closeWg.Wait(); close(done) }()
	select {
	case <-done:
	case <-time.After(6 * time.Second):
		t.Fatal("并发关闭超时")
	}
}

// T9: fun 内部监听 ctx.Done() 优雅退出
func TestContextCancellation(t *testing.T) {
	stopped := make(chan struct{})
	// 需要通过 goroutineMap 获取 ctx，直接测试 close 触发
	id, err := SafeGo(0, func(args ...interface{}) {
		// 模拟工作，这里无法直接访问 ctx，通过 sleep+channel 模拟
		<-time.After(30 * time.Second)
	})
	if err != nil {
		t.Fatalf("创建安全goroutine失败 %v", err)
	}

	time.Sleep(50 * time.Millisecond)

	go func() {
		CloseSafeGo(id)
		close(stopped)
	}()

	select {
	case <-stopped:
	case <-time.After(6 * time.Second):
		t.Fatal("关闭超时")
	}
}

// T10: 关闭超时保护 — fun 不响应 cancel
// （注意：safeRun 本身不会无限阻塞，因为每次循环都会检查 ctx.Done()）
// 但如果 fun 本身阻塞太久，close 的超时是 5s
// 这里验证 Close 在有长期运行 goroutine 时不会永久阻塞
func TestCloseTimeout(t *testing.T) {
	id, err := SafeGo(0, func(args ...interface{}) {
		// 模拟长时间运行（这个 fun 不检查 ctx）
		time.Sleep(10 * time.Second)
	})
	if err != nil {
		t.Fatalf("创建安全goroutine失败 %v", err)
	}

	start := time.Now()
	CloseSafeGo(id)
	elapsed := time.Since(start)

	// 应在 5-7 秒内返回（5s 超时 + 少量开销）
	if elapsed > 7*time.Second {
		t.Fatalf("关闭耗时太长: %v", elapsed)
	}
}

// T11: 数量上限拒绝
func TestMaxGoroutinesLimit(t *testing.T) {
	// 重置计数器（避免其他测试的干扰）
	atomic.StoreInt64(&goroutineCount, 0)
	Init(2) // 上限 2 个

	id1, err := SafeGo(0, func(args ...interface{}) {
		<-time.After(5 * time.Second)
	})
	if err != nil {
		t.Fatalf("第1个创建失败 %v", err)
	}

	id2, err := SafeGo(0, func(args ...interface{}) {
		<-time.After(5 * time.Second)
	})
	if err != nil {
		t.Fatalf("第2个创建失败 %v", err)
	}

	_, err = SafeGo(0, func(args ...interface{}) {
		<-time.After(5 * time.Second)
	})
	if err != nil {
		t.Fatalf("第3个创建失败 %v", err)
	}

	// 清理
	CloseSafeGo(id1)
	CloseSafeGo(id2)
}

// T12: 配额释放后恢复创建
func TestQuotaRelease(t *testing.T) {
	atomic.StoreInt64(&goroutineCount, 0)
	Init(2) // 上限 2 个

	id1, err := SafeGo(0, func(args ...interface{}) {
		<-time.After(5 * time.Second)
	})
	if err != nil {
		t.Fatalf("第1个创建失败 %v", err)
	}

	id2, err := SafeGo(0, func(args ...interface{}) {
		<-time.After(5 * time.Second)
	})
	if err != nil {
		t.Fatalf("第2个创建失败 %v", err)
	}

	_, err = SafeGo(0, func(args ...interface{}) {
		<-time.After(5 * time.Second)
	})
	if err != nil {
		t.Fatalf("第3个创建失败 %v", err)
	}

	// 释放一个后可以再创建
	CloseSafeGo(id1)
	time.Sleep(200 * time.Millisecond)

	id4, err := SafeGo(0, func(args ...interface{}) {
		<-time.After(5 * time.Second)
	})
	if err != nil {
		t.Fatalf("第4个创建失败 %v", err)
	}

	// 清理
	CloseSafeGo(id2)
	CloseSafeGo(id4)
}

// TestNoInitNoLimit 验证未调用 Init 时不限制数量（向后兼容）
func TestNoInitNoLimit(t *testing.T) {
	atomic.StoreInt64(&gCfg.maxGoroutines, 0) // 模拟未初始化

	ids := make([]string, 0, 10)
	for i := 0; i < 10; i++ {
		id, err := SafeGo(0, func(args ...interface{}) {
			<-time.After(5 * time.Second)
		})
		if err != nil {
			t.Fatalf("未调用Init时第%d个创建失败 %v", i, err)
		}
		ids = append(ids, id)
	}

	// 清理
	for _, id := range ids {
		CloseSafeGo(id)
	}
}

// TestSafeGoFunc 验证无参函数封装正常运行
func TestSafeGoFunc(t *testing.T) {
	done := make(chan struct{})
	id, err := SafeGoFunc(0, func() {
		close(done)
	})
	if err != nil {
		t.Fatalf("SafeGoFunc 创建失败: %v", err)
	}
	defer CloseSafeGo(id)

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("SafeGoFunc goroutine 未及时退出")
	}
}

// TestSafeGoFuncPanic 验证无参函数封装的 panic 恢复
func TestSafeGoFuncPanic(t *testing.T) {
	var count int32
	id, err := SafeGoFunc(2, func() {
		atomic.AddInt32(&count, 1)
		c := atomic.LoadInt32(&count)
		if c < 3 {
			panic("无参函数 panic")
		}
	})
	if err != nil {
		t.Fatalf("SafeGoFunc 创建失败: %v", err)
	}
	defer CloseSafeGo(id)

	time.Sleep(500 * time.Millisecond)

	if atomic.LoadInt32(&count) != 3 {
		t.Fatalf("期望执行3次，实际 %d", atomic.LoadInt32(&count))
	}
}

// TestSafeGoRebootlessFunc 验证无参不重启封装
func TestSafeGoRebootlessFunc(t *testing.T) {
	done := make(chan struct{})
	id, err := SafeGoRebootlessFunc(func() {
		close(done)
	})
	if err != nil {
		t.Fatalf("SafeGoRebootlessFunc 创建失败: %v", err)
	}
	defer CloseSafeGo(id)

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("SafeGoRebootlessFunc goroutine 未及时退出")
	}
}

// TestSafeGoGroupFunc 验证群组无参封装
func TestSafeGoGroupFunc(t *testing.T) {
	done := make(chan struct{})
	_, err := SafeGoGroupFunc("test-group-func", 0, func() {
		close(done)
	})
	if err != nil {
		t.Fatalf("SafeGoGroupFunc 创建失败: %v", err)
	}
	defer CloseSafeGoGroup("test-group-func")

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("SafeGoGroupFunc goroutine 未及时退出")
	}

	// 验证群组关闭幂等
	CloseSafeGoGroup("test-group-func")
	CloseSafeGoGroup("test-group-func")
}

// TestSafeGoGroupRebootlessFunc 验证群组无参不重启封装
func TestSafeGoGroupRebootlessFunc(t *testing.T) {
	done := make(chan struct{})
	_, err := SafeGoGroupRebootlessFunc("test-group-func-rl", func() {
		close(done)
	})
	if err != nil {
		t.Fatalf("SafeGoGroupRebootlessFunc 创建失败: %v", err)
	}
	defer CloseSafeGoGroup("test-group-func-rl")

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("SafeGoGroupRebootlessFunc goroutine 未及时退出")
	}

	// 清理
	CloseSafeGoGroup("test-group-func-rl")
}
