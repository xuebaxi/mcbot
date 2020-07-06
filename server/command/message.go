package command

import (
	"encoding/json"

	"github.com/xuebaxi/mcbot/server/nchan"
	"github.com/xuebaxi/mcbot/server/server"
)

func message([]string) string {
	str, err := json.Marshal(nchan.AllString(server.Server.Status["message"].(chan string)))
	if err != nil {
		return ""
	}
	return string(str)
}

func init() {
	Command["message"] = message
}
