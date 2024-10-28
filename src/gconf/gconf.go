package gconf

var (
	Conf = "./conf/server.json"
	// 默认Tcp配置
	TcpConf = map[string]any{
		"maxsize": 10000,       // 最大连接数
		"ip":      "127.0.0.1", // IP地址
		"port":    9001,        // 端口号
	}
	// 默认Http配置
	HttpConf = map[string]any{
		"ip":   "127.0.0.1", // IP地址
		"port": 9002,        // 端口号
	}
	// 默认WebSocket配置
	WebSocketConf = map[string]any{
		"ip":   "127.0.0.1", // IP地址
		"port": 9003,        // 端口号
	}
	// 默认Mysql数据库配置
	MysqlConf = map[string]any{
		"host":      "127.0.0.1",
		"port":      3306,
		"user":      "root",
		"password":  "111111",
		"database":  "test",
		"keepalive": true, // 是否保持连接
		"maxnum":    100,  // 最大连接数
		"maxidle":   20,   // 最大空闲连接数
	}
	// 默认Log配置
	LogConf = map[string]any{
		"path":   "./logs/",
		"level":  "error",
		"format": "%s【%s:%d】[%s]: %s", // time 【module:line】[level]: message
		"size":   1024 * 1024 * 100,   // 日志文件大小，单位为字节，默认为10MB
		"logfile": []map[string]string{
			{
				"level":    "info",
				"filename": "info",
			},
			{
				"level":    "warning",
				"filename": "warning",
			},
			{
				"level":    "error",
				"filename": "error",
			},
			{
				"level":    "crash",
				"filename": "crash",
			},
		},
	}
)
