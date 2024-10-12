package gmysql

import (
	"main/deps/glog"
	"time"
)

func main() {
	mc := map[string]any{
		"host":      "127.0.0.1",
		"port":      3306,
		"user":      "root",
		"password":  "111111",
		"database":  "test",
		"keepalive": true, // 是否保持连接
		"maxnum":    100,  // 最大连接数
		"maxidle":   20,   // 最大空闲连接数
	}
	err := Start(mc)
	if err != nil {
		glog.CrashExit(err.Error())
	}

	go ShowTables()
	go QueryTable("aa")
	go QueryTable("bb")
}

func ShowTables() {
	for {
		time.Sleep(time.Second * 5) // 每隔5s查询一次
		r, err := Query("show tables;")
		if err != nil {
			glog.Error(err.Error())
			continue
		}
		// defer r.Close()
		var tbn string
		for r.Next() {
			err = r.Scan(&tbn)
			if err != nil {
				glog.Error(err.Error())
			}
			glog.Info(tbn)
		}
	}
}

func QueryTable(tbn string) {
	for {
		time.Sleep(time.Second * 5) // 每隔3s查询一次
		r, err := Query("select * from " + tbn)
		if err != nil {
			glog.Error(err.Error())
			continue
		}
		// defer r.Close()
		var id, age int
		var name string
		for r.Next() {
			err = r.Scan(&id, &name, &age)
			if err != nil {
				glog.Error(err.Error())
			}
			glog.Info("%d %s %d", id, name, age)
		}
	}
}
