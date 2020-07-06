package command_test

import (
	"testing"

	"github.com/xuebaxi/mcbot/server/command"
	"github.com/xuebaxi/mcbot/server/nchan"
	"github.com/xuebaxi/mcbot/server/server"
)

func TestMessage(t *testing.T) {
	server.Server.Status["message"] = make(chan string, 10)
	nchan.WriteString(server.Server.Status["message"].(chan string), "test")
	result := command.Command["message"]([]string{})
	if result != "[\"test\"]" {
		panic("message 功能异常")
	}
}

func TestStatus(t *testing.T) {
	server.Server.Status["users"] = map[string]interface{}{
		"online":   int(0),
		"userlist": []string{},
	}
	server.Server.Status["users"].(map[string]interface{})["online"] =
		server.Server.Status["users"].(map[string]interface{})["online"].(int) + 1
	server.Server.Status["users"].(map[string]interface{})["userlist"] =
		append(server.Server.Status["users"].(map[string]interface{})["userlist"].([]string), "obj")
	result := command.Command["status"]([]string{})
	if result != "{\"online\":1,\"userlist\":[\"obj\"]}" {
		panic("status 功能异常")
	}
}
