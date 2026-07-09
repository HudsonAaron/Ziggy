package gsafego

import (
	"context"
	"fmt"
	"main/deps/glog"
	"sync"
	"sync/atomic"
	"time"
)

// Init 初始化安全 goroutine 管理器
//
//   - maxGoroutines: 可选参数，设置全局最大 goroutine 数量上限。
//     不传或传入 <=0 → 使用默认值 DefaultMaxGoroutines（100,000）
//
// 行为：
//  1. 设置 gCfg.maxGoroutines（原子写入）
//  2. 标记 gCfg.initialized = true
//  3. 可多次调用，后续调用仅更新上限值（不会重置计数器）
//
// 未调用 Init 时，maxGoroutines 默认为 0（不限制），以保持向后兼容。
func Init(maxGoroutines ...int) {
	limit := int64(DefaultMaxGoroutines)
	if len(maxGoroutines) > 0 && maxGoroutines[0] > 0 {
		limit = int64(maxGoroutines[0])
	}
	atomic.StoreInt64(&gCfg.maxGoroutines, limit)
	glog.Info("[gsafego] Init: maxGoroutines=%d", limit)
}

// Close 关闭所有安全 goroutine（先全部 cancel，再并行等待，避免串行阻塞）
func Close() {
	var wg sync.WaitGroup

	// 第一遍：收集所有 goroutine，发送取消信号（不等待）
	goroutineMap.Range(func(key, value any) bool {
		sg := value.(*safeGoroutine)
		sg.cancel()
		wg.Add(1)
		go func(s *safeGoroutine) {
			defer wg.Done()
			timer := time.NewTimer(5 * time.Second)
			defer timer.Stop()
			select {
			case <-s.done:
			case <-timer.C:
			}
		}(sg)
		return true
	})

	// 并行等待所有 goroutine 退出（最多 5s）
	wg.Wait()

	// 清理群组
	groupMap.Range(func(key, value any) bool {
		grp := value.(*safeGroup)
		groupMap.Delete(grp.id)
		return true
	})
	glog.Info("[gsafego] Close: 已关闭所有安全goroutine")
}

// SafeGo 创建独立安全 goroutine（单一模式）
//
//   - restartCount: 最大重启次数，<=0 表示不重启（panic 后直接退出）
//   - fun:          要执行的函数，通过 args 接收参数
//   - args:         传递给 fun 的可选参数
//
// 返回: goroutineID（唯一标识，用于后续 CloseSafeGoroutine）
//
//	若达到全局数量上限，返回空字符串 "" 并记录 Error 日志
func SafeGo(restartCount int, fun func()) (string, error) {
	if fun == nil {
		glog.Error("[gsafego] CreateSafeGoroutine: fun 不能为 nil")
		return "", fmt.Errorf("[gsafego] CreateSafeGoroutine: fun 不能为 nil")
	}
	if restartCount <= 0 {
		restartCount = 0 // 不重启
	}

	// 预留配额（先递增再检查，超限自动回滚，消除 TOCTOU 竞态）
	if !reserveGoroutineSlot() {
		glog.Error("[gsafego] 已达到最大goroutine数量上限(%d)，拒绝创建",
			atomic.LoadInt64(&gCfg.maxGoroutines))
		return "", fmt.Errorf("[gsafego] 已达到最大goroutine数量上限(%d)，拒绝创建", atomic.LoadInt64(&gCfg.maxGoroutines))
	}

	sg := &safeGoroutine{
		id:         generateID(),
		restartMax: restartCount,
		fun:        fun,
		done:       make(chan struct{}),
	}
	sg.ctx, sg.cancel = context.WithCancel(context.Background())

	goroutineMap.Store(sg.id, sg)
	go safeRun(sg)

	// glog.Info("[gsafego] 创建安全goroutine: id=%s, restartMax=%d", sg.id, sg.restartMax)
	return sg.id, nil
}

// SafeGoRebootless 创建不重启独立安全 goroutine（单一模式）
//
//   - fun:          要执行的函数，通过 args 接收参数
//   - args:         传递给 fun 的可选参数
//
// 返回: goroutineID（唯一标识，用于后续 CloseSafeGoroutine）
//
//	若达到全局数量上限，返回空字符串 "" 并记录 Error 日志
func SafeGoRebootless(fun func()) (string, error) {
	if fun == nil {
		glog.Error("[gsafego] CreateSafeGoroutine: fun 不能为 nil")
		return "", fmt.Errorf("[gsafego] CreateSafeGoroutine: fun 不能为 nil")
	}

	// 预留配额（先递增再检查，超限自动回滚，消除 TOCTOU 竞态）
	if !reserveGoroutineSlot() {
		glog.Error("[gsafego] 已达到最大goroutine数量上限(%d)，拒绝创建",
			atomic.LoadInt64(&gCfg.maxGoroutines))
		return "", fmt.Errorf("[gsafego] 已达到最大goroutine数量上限(%d)，拒绝创建", atomic.LoadInt64(&gCfg.maxGoroutines))
	}

	sg := &safeGoroutine{
		id:         generateID(),
		restartMax: 0,
		fun:        fun,
		done:       make(chan struct{}),
	}
	sg.ctx, sg.cancel = context.WithCancel(context.Background())

	goroutineMap.Store(sg.id, sg)
	go safeRun(sg)

	// glog.Info("[gsafego] 创建安全goroutine: id=%s, restartMax=%d", sg.id, sg.restartMax)
	return sg.id, nil
}

// CloseSafeGo 关闭独立安全 goroutine（单一模式）
//
//   - goroutineID: CreateSafeGoroutine 返回的 ID
//
// 幂等：多次调用同一个 ID 不会 panic
func CloseSafeGo(goroutineID string) {
	v, ok := goroutineMap.Load(goroutineID)
	if !ok {
		return // 幂等：已经不存在
	}
	sg := v.(*safeGoroutine)

	sg.cancel() // 发送取消信号

	// 等待完成，带超时保护（使用 NewTimer 避免 time.After 泄漏 timer）
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()
	select {
	case <-sg.done:
	case <-timer.C:
		glog.Warning("[gsafego] goroutine %s 关闭超时", goroutineID)
	}

	glog.Info("[gsafego] 关闭安全goroutine: id=%s", goroutineID)
}

// SafeGoGroup 创建群组并添加第一个安全 goroutine（群组模式）
//
//   - groupID:       群组 ID（若群组不存在则自动创建）
//   - restartCount:  最大重启次数（与单一模式语义一致）
//   - fun:           要执行的函数
//   - args:          传递给 fun 的可选参数
//
// 返回: goroutineID（该 goroutine 的唯一标识）
//
//	若达到全局数量上限，返回空字符串 "" 并记录 Error 日志
func SafeGoGroup(groupID string, restartCount int, fun func()) (string, error) {
	if fun == nil {
		glog.Error("[gsafego] CreateSafeGoroutineGroup: fun 不能为 nil")
		return "", fmt.Errorf("[gsafego] CreateSafeGoroutineGroup: fun 不能为 nil")
	}
	if restartCount <= 0 {
		restartCount = 0
	}

	// 预留配额（先递增再检查，超限自动回滚，消除 TOCTOU 竞态）
	if !reserveGoroutineSlot() {
		glog.Error("[gsafego] 已达到最大goroutine数量上限(%d)，拒绝创建",
			atomic.LoadInt64(&gCfg.maxGoroutines))
		return "", fmt.Errorf("[gsafego] 已达到最大goroutine数量上限(%d)，拒绝创建", atomic.LoadInt64(&gCfg.maxGoroutines))
	}

	// 确保群组存在
	grp, _ := groupMap.LoadOrStore(groupID, &safeGroup{
		id:      groupID,
		members: make(map[string]*safeGoroutine),
	})

	sg := &safeGoroutine{
		id:         generateID(),
		groupID:    groupID,
		restartMax: restartCount,
		fun:        fun,
		done:       make(chan struct{}),
	}
	sg.ctx, sg.cancel = context.WithCancel(context.Background())

	goroutineMap.Store(sg.id, sg)

	// 加入群组
	g := grp.(*safeGroup)
	g.mu.Lock()
	g.members[sg.id] = sg
	g.mu.Unlock()

	go safeRun(sg)

	// glog.Info("[gsafego] 创建群组安全goroutine: id=%s, groupID=%s, restartMax=%d",
	// 	sg.id, groupID, sg.restartMax)
	return sg.id, nil
}

// SafeGoGroupRebootless 创建不重启的群组并添加第一个安全 goroutine（群组模式）
//
//   - groupID:       群组 ID（若群组不存在则自动创建）
//   - fun:           要执行的函数
//   - args:          传递给 fun 的可选参数
//
// 返回: goroutineID（该 goroutine 的唯一标识）
//
//	若达到全局数量上限，返回空字符串 "" 并记录 Error 日志
func SafeGoGroupRebootless(groupID string, fun func()) (string, error) {
	if fun == nil {
		glog.Error("[gsafego] CreateSafeGoroutineGroup: fun 不能为 nil")
		return "", fmt.Errorf("[gsafego] CreateSafeGoroutineGroup: fun 不能为 nil")
	}

	// 预留配额（先递增再检查，超限自动回滚，消除 TOCTOU 竞态）
	if !reserveGoroutineSlot() {
		glog.Error("[gsafego] 已达到最大goroutine数量上限(%d)，拒绝创建",
			atomic.LoadInt64(&gCfg.maxGoroutines))
		return "", fmt.Errorf("[gsafego] 已达到最大goroutine数量上限(%d)，拒绝创建", atomic.LoadInt64(&gCfg.maxGoroutines))
	}

	// 确保群组存在
	grp, _ := groupMap.LoadOrStore(groupID, &safeGroup{
		id:      groupID,
		members: make(map[string]*safeGoroutine),
	})

	sg := &safeGoroutine{
		id:         generateID(),
		groupID:    groupID,
		restartMax: 0,
		fun:        fun,
		done:       make(chan struct{}),
	}
	sg.ctx, sg.cancel = context.WithCancel(context.Background())

	goroutineMap.Store(sg.id, sg)

	// 加入群组
	g := grp.(*safeGroup)
	g.mu.Lock()
	g.members[sg.id] = sg
	g.mu.Unlock()

	go safeRun(sg)
	return sg.id, nil
}

// CloseSafeGoGroup 关闭整个群组（群组模式）
//
//   - groupID: 群组 ID
//
// 幂等：多次调用同一个 groupID 不会 panic
func CloseSafeGoGroup(groupID string) {
	v, ok := groupMap.Load(groupID)
	if !ok {
		return // 幂等：已经不存在
	}
	grp := v.(*safeGroup)

	// 获取成员快照
	grp.mu.RLock()
	memberIDs := make([]string, 0, len(grp.members))
	for id := range grp.members {
		memberIDs = append(memberIDs, id)
	}
	grp.mu.RUnlock()

	// 逐个关闭（并发，先 cancel 再 wait，避免串行阻塞）
	var wg sync.WaitGroup
	for _, id := range memberIDs {
		v, ok := goroutineMap.Load(id)
		if !ok {
			continue
		}
		sg := v.(*safeGoroutine)
		sg.cancel()
		wg.Add(1)
		go func(s *safeGoroutine) {
			defer wg.Done()
			timer := time.NewTimer(5 * time.Second)
			defer timer.Stop()
			select {
			case <-s.done:
			case <-timer.C:
			}
		}(sg)
	}
	wg.Wait()

	groupMap.Delete(groupID)
	glog.Info("[gsafego] 关闭安全goroutine群组: groupID=%s, 关闭成员数=%d",
		groupID, len(memberIDs))
}
