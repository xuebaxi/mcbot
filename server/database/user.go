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

type user struct {
	db        *sql.DB
	tableName string
}

//参数为sql命令,防止数据库锁死.
func (db *user) Read(sqlcmd string) ([]map[string]string, error) {
	access.Lock()
	defer access.Unlock()
	sqlcmd = strings.Replace(sqlcmd, "$tablename", db.tableName, -1)
	var (
		id                  int
		userid, permissions string
	)
	rows, err := db.db.Query(sqlcmd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result = []map[string]string{}
	for rows.Next() {
		err := rows.Scan(&id, &userid, &permissions)
		if err != nil {
			return nil, err
		}
		result = append(result, map[string]string{
			"id":          strconv.Itoa(id),
			"userid":      userid,
			"permissions": permissions,
		})
	}
	return result, nil
}

func (db *user) Write(m map[string]string) error {
	access.Lock()
	defer access.Unlock()
	var (
		userid, permissions string
		ok                  bool
		err                 error
	)
	if userid, ok = m["userid"]; !ok {
		log.Errorln("数据错误,需要用户ID.")
		return fmt.Errorf("数据错误,需要用户ID")
	}
	if permissions, ok = m["permissions"]; !ok {
		log.Errorln("数据错误,需要权限.")
		return fmt.Errorf("数据错误,需要权限")
	}

	stmt, err := db.db.Prepare(fmt.Sprintf("INSERT INTO %s(userid, permissions) values(?,?)", db.tableName))
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userid, permissions)
	if err != nil {
		return err
	}
	return nil
}

func (db *user) init(sqldb *sql.DB) error {
	access.Lock()
	defer access.Unlock()
	db.db = sqldb
	db.tableName = "users"
	_, err := db.db.Exec(fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS  %s (
		id INTEGER PRIMARY KEY, 
		userid TEXT, 
		permissions	 TEXT
		)
	`, db.tableName))
	if err != nil {
		return err
	}
	return nil
}
