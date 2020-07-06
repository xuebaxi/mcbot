package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	//加载数据库驱动
	_ "github.com/mattn/go-sqlite3"

	log "github.com/sirupsen/logrus"
)

type even struct {
	db        *sql.DB
	tableName string
}

//参数为sql命令,防止数据库锁死.
func (db *even) Read(sqlcmd string) ([]map[string]string, error) {
	access.Lock()
	defer access.Unlock()
	sqlcmd = strings.Replace(sqlcmd, "$tablename", db.tableName, -1)
	var (
		id                                   int
		dname, dtime, objtype, obj, dcontent string
	)
	rows, err := db.db.Query(sqlcmd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result = []map[string]string{}
	for rows.Next() {
		err := rows.Scan(&id, &dname, &dtime, &objtype, &obj, &dcontent)
		if err != nil {
			return nil, err
		}
		result = append(result, map[string]string{
			"id":      strconv.Itoa(id),
			"name":    dname,
			"time":    dtime,
			"objtype": objtype,
			"obj":     obj,
			"content": dcontent,
		})
	}
	return result, nil
}

func (db *even) Write(m map[string]string) error {
	access.Lock()
	defer access.Unlock()
	var (
		dname, dtime, dcontent, dobj, dobjtype string
		ok                                     bool
		err                                    error
	)
	if dname, ok = m["name"]; !ok {
		log.Errorln("数据错误,需要事件名.")
		return fmt.Errorf("数据错误,需要事件名")
	}
	if dtime, ok = m["time"]; !ok {
		log.Errorln("数据错误,需要事件发生时间.")
		return fmt.Errorf("数据错误,需要事件发生时间")
	}
	if dcontent, ok = m["content"]; !ok {
		log.Errorln("数据错误,需要事件内容.")
		return fmt.Errorf("数据错误,需要事件内容")
	}
	if dobj, ok = m["obj"]; !ok {
		log.Errorln("数据错误,需要事件对象.")
		return fmt.Errorf("数据错误,需要事件对象")
	}
	if dobjtype, ok = m["objtype"]; !ok {
		log.Errorln("数据错误,需要事件对象类型.")
		return fmt.Errorf("数据错误,需要事件对象类型")
	}
	stmt, err := db.db.Prepare(fmt.Sprintf("INSERT INTO %s(name, time,objtype,obj,content) values(?,?,?,?,?)", db.tableName))
	if err != nil {
		return err
	}
	_, err = stmt.Exec(dname, dtime, dobjtype, dobj, dcontent)
	if err != nil {
		return err
	}
	return nil
}

func (db *even) init(sqldb *sql.DB) error {
	access.Lock()
	defer access.Unlock()
	db.db = sqldb
	db.tableName = "evens"
	_, err := db.db.Exec(fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS  %s (
		id INTEGER PRIMARY KEY, 
		name TEXT, 
		time TEXT,
		objtype TEXT,
		obj  TEXT,
		content TEXT
		)
	`, db.tableName))
	if err != nil {
		return err
	}
	return nil
}
