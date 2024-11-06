package gconf

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// 服务配置结构体
type ServerConf struct {
	TcpConf       map[string]any `json:"TcpConf"`       // Tcp配置
	HttpConf      map[string]any `json:"HttpConf"`      // Http配置
	WebSocketConf map[string]any `json:"WebSocketConf"` // WebSocket配置
	MysqlConf     map[string]any `json:"MysqlConf"`     // Mysql配置
	LogConf       map[string]any `json:"LogConf"`       // 日志配置
}

// 读取服务配置
func LoadConf() error {
	file, err := os.Open(Conf)
	if err != nil {
		return err
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	SrvConf := ServerConf{}
	err = json.Unmarshal(bytes, &SrvConf)
	if err != nil {
		return err
	}
	fmt.Println("LoadConf Success")
	TcpConf = SrvConf.TcpConf
	HttpConf = SrvConf.HttpConf
	WebSocketConf = SrvConf.WebSocketConf
	MysqlConf = SrvConf.MysqlConf
	LogConf = SrvConf.LogConf
	return nil
}
