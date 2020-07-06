package main

import (
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/xuebaxi/mcbot/server/app"
	_ "github.com/xuebaxi/mcbot/server/database"
	"github.com/xuebaxi/mcbot/server/load"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: false,
	})
	var logobj = log.New()
	logobj.SetLevel(log.DebugLevel)
	go func() {
		load.Loader(logobj)
	}()
	router := gin.New()
	router.Static("/assets", "./assets")
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	app.RegisterRoutes(router)
	gin.SetMode(gin.DebugMode)
	router.Run(":8080")
}
