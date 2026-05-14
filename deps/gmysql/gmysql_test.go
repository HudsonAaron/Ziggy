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

// 存储过程用例
func PrepareSQL() {
	// 首先删除旧的存储过程（如果存在）
	_, err := gmysql.Exec("DROP PROCEDURE IF EXISTS selectRole")
	if err != nil {
		glog.Error("删除旧存储过程失败: %s", err.Error())
	}
	// 不使用prepared statement创建存储过程，直接使用Exec
	_, err = Exec("create procedure if not exists selectRole() begin select roleId,nickname,mpId from role; end")
	if err != nil {
		glog.Error(err.Error())
		return
	}

	// 调用存储过程时使用Query而不是CallPrepare
	rows, err := Query("call selectRole()")
	if err != nil {
		glog.Error(err.Error())
		return
	}
	defer rows.Close()
	// 遍历结果集
	for rows.Next() {
		var roleId int
		var nickname string
		var MpId int
		err := rows.Scan(&roleId, &nickname, &MpId)
		if err != nil {
			glog.Error(err.Error())
			return
		}
		glog.Info("ID: %d, Name: %s, MpId: %d", roleId, nickname, MpId)
	}

	// 处理查询结果
	glog.Info("存储过程调用成功")
	// 这里可以添加对结果集的处理代码
}
