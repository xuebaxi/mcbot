package command

import (
	"fmt"

	"github.com/xuebaxi/mcbot/server/server"
)

func run(s []string) string {
	var texts string
	for _, text := range s[1:] {
		texts += text
		texts += " "
	}
	texts = texts[:len(texts)-1]
	server.Server.Run(fmt.Sprintf("/%s", texts))
	return ""
}

func init() {
	Admin["run"] = run
}
