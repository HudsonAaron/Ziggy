package gmysql

import (
	"database/sql"
	"main/deps/glog"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Version = "1.0.0" // 版本号
	DBConn  *sql.DB
)

// 连接数据库连接
func Start(mysqlConf map[string]any) error {
	mc, err := GetMysqlConf(mysqlConf)
	if err != nil {
		return err
	}
	dataSrcName := mc.GetDBSrcName()
	db, err := sql.Open("mysql", dataSrcName)
	if err != nil {
		return err
	}
	DBConn = db
	// 设置最大连接数
	db.SetMaxOpenConns(mc.MaxNum)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(mc.MaxIdle)
	// 设置保持活跃连接
	if mc.Keepalive {
		db.SetConnMaxLifetime(30 * time.Second)
	}
	glog.Info("DB Server start up")
	return nil
}

// 关闭数据库连接
func Stop() {
	if DBConn != nil {
		DBConn.Close()
	}
	glog.Crash("DB Server stop")
}

// 执行SQL语句
func Exec(sql string) (sql.Result, error) {
	return DBConn.Exec(sql)
}

// 执行SQL语句，使用事务
func ExecTx(sqls []string) error {
	tx, err := DBConn.Begin()
	if err != nil {
		return err
	}
	for _, sql := range sqls {
		_, err := tx.Exec(sql)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

// 查询SQL语句
func Query(sql string) (*sql.Rows, error) {
	return DBConn.Query(sql)
}
