package command

import (
	"encoding/json"

	"github.com/xuebaxi/mcbot/server/server"
)

func status([]string) string {
	str, err := json.Marshal(server.Server.Status["users"])
	if err != nil {
		return ""
	}
	return string(str)
}

func init() {
	Command["status"] = status
}
