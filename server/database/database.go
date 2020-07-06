package database

import (
	"database/sql"
	"os"
	"sync"

	//加载数据库驱动
	_ "github.com/mattn/go-sqlite3"

	log "github.com/sirupsen/logrus"
)

//DataBases 是用于控制数据库的单例.
var access sync.RWMutex
var DataBases map[string]database

type database interface {
	Read(string) ([]map[string]string, error)
	Write(map[string]string) error
	init(*sql.DB) error
}

//单例创建数据库
func newdatabase(path string) map[string]database {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Errorln("数据库初始化异常,打开数据库错误.")
		os.Exit(1)
	}
	var result = map[string]database{
		"even": new(even),
		"user": new(user),
	}
	for _, databaseObj := range result {
		err := databaseObj.init(db)
		if err != nil {
			log.Errorln("数据库创建失败")
			os.Exit(1)
		}
	}
	log.Infoln("初始化数据库完成")
	return result
}
func init() {
	if DataBases == nil {
		DataBases = newdatabase("mcbots.db")
	}
}
