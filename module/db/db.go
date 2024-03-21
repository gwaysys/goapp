package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gwaylib/conf"
	"github.com/gwaylib/qsql"
)

func init() {
	qsql.RegCacheWithIni(conf.RootDir() + "/etc/db.cfg")
}

func GetCache(section string) *qsql.DB {
	return qsql.GetCache(section)
}

func HasCache(section string) (*qsql.DB, error) {
	return qsql.HasCache(section)
}

// 当使用了Cache，在程序退出时可调用qsql.CloseCache进行正常关闭数据库连接
func CloseCache() {
	qsql.CloseCache()
}
