package glog

import (
	"log"
	"main/deps/gutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// 日志结构体
type Logger struct {
	file       *os.File    // 日志文件
	filePath   string      // 日志文件路径
	filename   string      // 日志名称
	path       string      // 日志路径
	level      string      // 日志等级
	maxSize    int64       // 日志文件最大大小
	createTime int64       // 日志创建时间
	msgChan    chan string // 日志消息通道
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
func NewLogger(logLv string, logPath string, logName string, maxSize int64) *Logger {
	msgChan := make(chan string) // 创建消息通道
	lg := &Logger{
		level:    logLv,
		filename: logName,
		path:     logPath,
		maxSize:  maxSize,
		msgChan:  msgChan,
	}
	return lg
}

// 打开日志，日志文件不存在，则创建并打开
func (lg *Logger) OpenFile() {
	filePath := GetLogPath(lg.path, lg.filename) // 拼接日志路径
	// 如果文件已存在，则接下着写文件
	if _, err := os.Stat(filePath); err == nil {
		lg.file, err = os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0777)
		if err != nil {
			log.Fatal(err)
		}
	} else if lg.file == nil {
		lg.file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
	lg.filePath = filePath
	lg.createTime = gutil.Timestamp()
}

// 启动日志输出
func (lg *Logger) Start() {
	for {
		msg := <-lg.msgChan
		if lg.CheckLogRebuild() {
			lg.Stop()     // 关闭文件
			lg.OpenFile() // 重新打开文件
		}
		lg.file.WriteString(msg + "\n")
	}
}

// 关闭日志文件
func (lg *Logger) Stop() {
	if lg != nil && lg.file != nil {
		lg.file.Close()
		lg.BatchRenameLog() // 批量修改日志文件名
		lg.file = nil
		lg.createTime = 0
		lg.filePath = ""
	}
}

// 拼接日志路径
func GetLogPath(logPath string, logName string) string {
	dateTime := gutil.FormatTimeByLayout("2006-01-02")
	return filepath.Join(logPath, logName+"_"+dateTime+".log")
}

// 检测日志文件状态
func (lg *Logger) CheckLogStatus() {
	if lg.CheckLogRebuild() {
		lg.Stop() // 关闭文件
	}
	if lg.file == nil { // 文件不存在，则打开文件
		lg.OpenFile()
	}
}

// 检测日志文件是否需要重建
func (lg *Logger) CheckLogRebuild() bool {
	if lg.file != nil {
		fileInfo, err := os.Stat(lg.filePath)
		if err != nil {
			log.Fatal(err)
			return false
		}
		// 文件大小超过限制，则关闭文件，重新创建新的日志文件
		IsOverSize := fileInfo.Size() > lg.maxSize
		// 判断当前时间是否跨越天，如果是，则关闭文件，重新创建新的日志文件
		IsSameDay := gutil.IsSameDay(gutil.Timestamp(), lg.createTime)
		if IsOverSize || !IsSameDay {
			return true
		}
		return false
	}
	return true
}

// 批量修改当前旧日志文件名
func (lg *Logger) BatchRenameLog() {
	if !lg.CheckLogRebuild() {
		return
	}
	// 获取当前日志文件所在路径下的所有相同名字的日志文件列表
	basePath := filepath.Dir(lg.filePath)
	files, err := os.ReadDir(basePath)
	if err != nil {
		log.Fatal(err)
	}
	// 去掉filePath中的路径以及后缀名
	filename := filepath.Base(lg.filePath)
	fName := filename[:len(filename)-len(filepath.Ext(filename))] // 无后缀
	logList := make([]string, 0)
	for _, file := range files {
		// 如果 file.Name() 包含了 fName，则添加到 logList 列表中
		if strings.Contains(file.Name(), fName) {
			logList = append(logList, file.Name())
		}
	}
	// 按照logList列表中的文件创建时间进行排序
	sortLogList := SortLogList(logList, filename)
	for _, logName := range sortLogList {
		num := "0"
		// 截取文件后缀后面的数字，没有数字，则默认为0
		if logName == filename {
			num = "0"
		} else {
			num = logName[len(filename)+1:]
			// num+1后，转为字符串
			idx, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			num = strconv.Itoa(idx + 1)
		}
		newName := filename + "." + num
		err := os.Rename(filepath.Join(basePath, logName), filepath.Join(basePath, newName))
		if err != nil {
			log.Fatal(err)
		}
	}
}

// 文件名后缀的序号越大，则越先开始处理
func SortLogList(logList []string, filename string) []string {
	for i := 0; i < len(logList); i++ {
		for j := i + 1; j < len(logList); j++ {
			if logList[j] == filename || (logList[i] == filename && logList[j] == filename) {
				continue
			}
			if logList[i] == filename && logList[j] != filename {
				logList[i], logList[j] = logList[j], logList[i]
				continue
			}
			if logList[i][len(filename)+1:] < logList[j][len(filename)+1:] {
				logList[i], logList[j] = logList[j], logList[i]
			}
		}
	}
	return logList
}
