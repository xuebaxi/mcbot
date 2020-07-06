package command

import (
	"github.com/xuebaxi/mcbot/server/server"
)

func debug(s []string) string {
	var texts string
	for _, text := range s[1:] {
		texts += text
		texts += " "
	}
	texts = texts[:len(texts)-1]
	if texts == "true" {
		server.Server.Status["debug"] = true
		return "已开启debug模式"
	} else if texts == "false" {
		server.Server.Status["debug"] = false
		return "已关闭debug模式"
	} else {
		return "无效输入"
	}
}

func init() {
	Admin["debug"] = debug
}
