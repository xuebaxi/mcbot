package database_test

import (
	"fmt"
	"runtime"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/xuebaxi/mcbot/server/database"
)

//测试数据库多线程读写
func TestDB(t *testing.T) {
	for _, obj := range database.DataBases {
		err := obj.Write(
			map[string]string{
				"name":        "test",
				"time":        "0",
				"eventype":    "test",
				"content":     "test",
				"objtype":     "server",
				"obj":         "mcserver",
				"userid":      "1",
				"permissions": "admin",
			},
		)
		if err != nil {
			log.Panic(err)
			return
		}
	}
	runtime.GOMAXPROCS(3) // 最多使用3个核,实现并行.
	w := make(chan error)
	r1 := make(chan error)
	r2 := make(chan error)
	go func() {
		for _, obj := range database.DataBases {
			result, err := obj.Read(
				`SELECT * FROM $tablename WHERE id=1`,
			)
			if err != nil {
				r1 <- err
				return
			}
			if !(result[0]["time"] == "0" || result[0]["userid"] == "1") {
				r1 <- fmt.Errorf("read database error")
				return
			}
		}

		r1 <- nil
	}()
	go func() {
		for _, obj := range database.DataBases {
			result, err := obj.Read(
				`SELECT * FROM $tablename WHERE id=1`,
			)
			if err != nil {
				r2 <- err
				return
			}
			if !(result[0]["time"] == "0" || result[0]["userid"] == "1") {
				r2 <- fmt.Errorf("read database error")
				return
			}
		}

		r2 <- nil
	}()
	go func() {
		for _, obj := range database.DataBases {
			err := obj.Write(
				map[string]string{
					"name":        "test",
					"time":        "0",
					"eventype":    "test",
					"content":     "test",
					"objtype":     "server",
					"obj":         "mcserver",
					"userid":      "1",
					"permissions": "admin",
				},
			)
			if err != nil {
				w <- err
				return
			}
		}

		w <- nil
	}()

	if err := <-w; err != nil {
		log.Panic(err)
	}
	if err := <-r1; err != nil {
		log.Panic(err)
	}
	if err := <-r2; err != nil {
		log.Panic(err)
	}
}
