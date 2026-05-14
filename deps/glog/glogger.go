package glog

import (
	"fmt"
	"log"
	"main/deps/gutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// 日志结构体（改进：添加锁和关闭控制）
type Logger struct {
	file       *os.File      // 日志文件
	filePath   string        // 日志文件路径
	filename   string        // 日志名称
	path       string        // 日志路径
	level      string        // 日志等级
	maxSize    int64         // 日志文件最大大小
	createTime int64         // 日志创建时间
	msgChan    chan string   // 日志消息通道
	mu         sync.Mutex    // 保护文件操作
	done       chan struct{} // 停止信号
}

// 日志等级
const (
	INFO    = 1 // info
	WARNING = 2 // warning
	ERROR   = 3 // error
	CRASH   = 4 // crash
)

// 确认日志路径，并创建
func MakeSureDir(filePath string) error {
	// 判断路径是否存在，不存在，则创建路径
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

// 创建指定类型日志文件
func newLogger(logLv string, logPath string, logName string, maxSize int64) *Logger {
	msgChan := make(chan string, 1000) // 有缓冲，避免阻塞
	done := make(chan struct{})
	lg := &Logger{
		level:    logLv,
		filename: logName,
		path:     logPath,
		maxSize:  maxSize,
		msgChan:  msgChan,
		done:     done,
	}
	return lg
}

// openFileLocked 打开或创建日志文件；调用方必须已持有 lg.mu。
// handleMsg 在持锁路径里会调用本函数，不可再调 openFile()（否则会二次 Lock 死锁）。
func (lg *Logger) openFileLocked() {
	filePath := GetLogPath(lg.path, lg.filename)
	if _, err := os.Stat(filePath); err == nil {
		// 追加模式
		f, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("open log file failed: %v", err)
			return
		}
		lg.file = f
	} else {
		// 创建新文件
		f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			fmt.Printf("create log file failed: %v", err)
			return
		}
		lg.file = f
	}
	lg.filePath = filePath
	lg.createTime = gutil.Timestamp()
}

// 打开日志，日志文件不存在，则创建并打开
func (lg *Logger) openFile() {
	lg.mu.Lock()
	defer lg.mu.Unlock()
	lg.openFileLocked()
}

// 启动日志输出（改进：支持优雅关闭，避免阻塞）
func (lg *Logger) Start() {
	lg.openFile() // 初始打开文件
	go lg.handleMsg()
}

// 关闭日志文件（改进：发送停止信号，支持优雅退出）
func (lg *Logger) stop() {
	if lg == nil || lg.done == nil {
		return
	}
	close(lg.done) // 通知Start goroutine退出
	// 等待一小段时间让Start goroutine完成最后写入
	time.Sleep(100 * time.Millisecond)
}

func (lg *Logger) handleMsg() {
	for {
		select {
		case msg := <-lg.msgChan:
			lg.mu.Lock()
			rebuildReason := lg.checkRebuildReason()
			if rebuildReason != "" {
				// 如果是因为文件大小超限，批量重命名当天旧文件（需在关闭前执行，此时 filePath 仍有效）
				if rebuildReason == "oversize" {
					lg.batchRenameLogs()
				}
				lg.closeFileLocked()
				lg.openFileLocked()
			}
			if lg.file != nil {
				_, _ = lg.file.WriteString(msg + "\n")
			}
			lg.mu.Unlock()
		case <-lg.done:
			lg.mu.Lock()
			lg.closeFileLocked()
			lg.mu.Unlock()
			return
		}
	}
}

// 拼接日志路径
func GetLogPath(logPath string, logName string) string {
	dateTime := gutil.FormatTimeByLayout("2006-01-02")
	dateTime := gutil.FormatTimeByLayout("2006-01-02")
	return filepath.Join(logPath, logName+"_"+dateTime+".log")
}

// 检测日志文件是否需要重建，返回重建原因（空字符串=不需要重建）
func (lg *Logger) checkRebuildReason() string {
	if lg.file == nil || lg.filePath == "" {
		return "nil"
	}
	fileInfo, err := os.Stat(lg.filePath)
	if err != nil {
		fmt.Printf("stat log file error: %v", err)
		return "stat_error"
	}
	// 文件大小超过限制
	if fileInfo.Size() > lg.maxSize {
		return "oversize"
	}
	// 跨天
	if !gutil.IsSameDay(gutil.Timestamp(), lg.createTime) {
		return "new_day"
	}
	return ""
}

// 关闭文件（带锁保护）
func (lg *Logger) closeFileLocked() {
	if lg.file != nil {
		lg.file.Close()
		lg.file = nil
		lg.filePath = ""
		lg.createTime = 0
	}
}

// 文件大小超限时，批量重命名当天所有同名日志文件
// 规则：按文件修改时间排序，时间越早的文件名字后带的数字越大
func (lg *Logger) batchRenameLogs() {
	if lg.filePath == "" {
		return
	}
	basePath := filepath.Dir(lg.filePath)
	filename := filepath.Base(lg.filePath)
	// 去掉后缀名
	prefix := strings.TrimSuffix(filename, filepath.Ext(filename))
	// 去掉路径
	filename = filepath.Base(filename)
	// 读取目录下的所有文件
	files, err := os.ReadDir(basePath)
	if err != nil {
		fmt.Printf("read dir error: %v", err)
		return
	}
	// 收集当天同名超限文件（排除当前主文件）
	type oldFile struct {
		name    string
		modTime time.Time
	}
	var oldFiles []oldFile
	for _, file := range files {
		name := file.Name()
		if !strings.HasPrefix(name, prefix) || !strings.Contains(name, ".log") {
			continue
		}
		info, err := file.Info()
		if err != nil {
			continue
		}
		oldFiles = append(oldFiles, oldFile{
			name:    name,
			modTime: info.ModTime(),
		})
	}
	if len(oldFiles) == 0 {
		return
	}

	// 按修改时间降序排序（最新的在前，最早的在后）
	for i := 0; i < len(oldFiles)-1; i++ {
		for j := i + 1; j < len(oldFiles); j++ {
			if oldFiles[i].modTime.Before(oldFiles[j].modTime) {
				oldFiles[i], oldFiles[j] = oldFiles[j], oldFiles[i]
			}
		}
	}

	// 最新的文件序号为1，越早创建的文件序号越大（倒序）
	for i := len(oldFiles) - 1; i >= 0; i-- {
		of := oldFiles[i]
		newName := fmt.Sprintf("%s.%d", filename, i+1)
		err = os.Rename(
			filepath.Join(basePath, of.name),
			filepath.Join(basePath, newName),
		)
		if err != nil {
			log.Printf("rename log file %s failed: %v", of.name, err)
		}
	}
}
