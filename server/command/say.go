package command

import (
	"fmt"

	"github.com/xuebaxi/mcbot/server/server"
)

func say(s []string) string {
	var texts string
	for _, text := range s[1:] {
		texts += text
		texts += " "
	}
	server.Server.Run(fmt.Sprintf("/say %s", texts))
	return ""
}

func init() {
	Command["say"] = say
}
