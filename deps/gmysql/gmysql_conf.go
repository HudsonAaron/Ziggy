package gmysql

import (
	"fmt"
)

// 数据库配置
type GMysqlConf struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Database  string `json:"database"`
	Keepalive bool   `json:"keepalive"` // 是否保持连接
	MaxNum    int    `json:"maxnum"`    // 最大连接数
	MaxIdle   int    `json:"maxidle"`   // 最大空闲连接数
}

// 获取数据库配置
func GetMysqlConf(mysqlConf map[string]any) (*GMysqlConf, error) {
	port := 3306 // 默认端口
	if mysqlConf["port"] != nil {
		if _port, ok := mysqlConf["port"].(int); ok {
			port = _port
		} else if _port, ok := mysqlConf["port"].(float64); ok {
			port = int(_port)
		} else {
			return nil, fmt.Errorf("mysql port type error")
		}
	}
	keepalive := false // 默认不保持连接
	if mysqlConf["keepalive"] != nil {
		if _keepalive, ok := mysqlConf["keepalive"].(bool); ok {
			keepalive = _keepalive
		} else {
			return nil, fmt.Errorf("mysql keepalive type error")
		}
	}
	maxnum := 10 // 默认最大连接数
	if mysqlConf["maxnum"] != nil {
		if _maxnum, ok := mysqlConf["maxnum"].(int); ok {
			maxnum = _maxnum
		} else if _maxnum, ok := mysqlConf["maxnum"].(float64); ok {
			maxnum = int(_maxnum)
		} else {
			return nil, fmt.Errorf("mysql maxnum type error")
		}
	}
	maxidle := 0 // 默认最大空闲连接数
	if mysqlConf["maxidle"] != nil {
		if _maxidle, ok := mysqlConf["maxidle"].(int); ok {
			maxidle = _maxidle
		} else if _maxidle, ok := mysqlConf["maxidle"].(float64); ok {
			maxidle = int(_maxidle)
		} else {
			return nil, fmt.Errorf("mysql maxidle type error")
		}
	}
	mc := &GMysqlConf{
		Host:      mysqlConf["host"].(string),
		Port:      port,
		User:      mysqlConf["user"].(string),
		Password:  mysqlConf["password"].(string),
		Database:  mysqlConf["database"].(string),
		Keepalive: keepalive,
		MaxNum:    maxnum,
		MaxIdle:   maxidle,
	}
	return mc, nil
}

// 获取数据库连接名
func (gmc *GMysqlConf) GetDBSrcName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		gmc.User,
		gmc.Password,
		gmc.Host,
		gmc.Port,
		gmc.Database,
	)
}
