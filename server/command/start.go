package command

import (
	"github.com/xuebaxi/mcbot/server/server"
)

func start(s []string) string {
	server.Server.Start()
	return ""
}

func init() {
	Admin["start"] = start
}
