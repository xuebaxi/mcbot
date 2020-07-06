package load

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/xuebaxi/mcbot/server/database"
	"github.com/xuebaxi/mcbot/server/nchan"
	"github.com/xuebaxi/mcbot/server/server"
)

type even struct {
	log *logrus.Logger
}

func (e *even) init(log *logrus.Logger) {
	server.Server.Status["users"] = map[string]interface{}{
		"online":   int(0),
		"userlist": []string{},
	}
	server.Server.Status["message"] = make(chan string, 10)
	server.Server.Status["running"] = bool(false)
	server.Server.Status["debug"] = bool(false)
	e.log = log
}

func (e *even) Save(text string) {
	e.log.Info(text)
	info := strings.Split(text, " ")
	if info[2] != "thread/INFO]:" || len(info) < 5 {
		return
	}
	if info[3] == "[Server]" {
		return
	}
	if server.Server.Status["debug"].(bool) {
		nchan.CoverString(server.Server.Status["message"].(chan string), text[len(info[0])+len(info[1])+len(info[2])+3:])
	}
	if info[3][0] == '<' && info[3][len(info[3])-1] == '>' {
		var content string = ""
		for i := 4; i < len(info); i++ {
			content += info[i]
			content += " "
		}
		e.saveSay(strconv.FormatInt(time.Now().Unix(), 10), info[3], content)
		return
	}
	switch info[3] {
	case "Done":
		/*Done (48.723s)! For help, type "help"*/
		if len(info) > 5 && info[5] == "For" {
			e.saveStart(strconv.FormatInt(time.Now().Unix(), 10))
			server.Server.Status["running"] = true
			e.log.Info("服务器启动成功")
			return
		}
	case "ThreadedAnvilChunkStorage":
		/*ThreadedAnvilChunkStorage (DIM1): All chunks are saved*/
		if info[4] == "(DIM1):" {
			e.saveStop(strconv.FormatInt(time.Now().Unix(), 10))
			server.Server.Status["running"] = false
			e.log.Info("服务器关闭")
			server.Server.Server.Wait()
			return
		}
	}
	if info[4] == "logged" {
		e.saveLogin(strconv.FormatInt(time.Now().Unix(), 10), info[3])
		return
	} else if info[4] == "left" {
		e.saveLogout(strconv.FormatInt(time.Now().Unix(), 10), info[3])
		return
	}
	if server.Server.Status["running"].(bool) && !(server.Server.Status["debug"].(bool)) {
		nchan.CoverString(server.Server.Status["message"].(chan string), text[len(info[0])+len(info[1])+len(info[2])+3:])
	}
}
func (e *even) saveStart(t string) {
	nchan.CoverString(server.Server.Status["message"].(chan string), "服务器已启动")
	err := database.DataBases["even"].Write(
		map[string]string{
			"name":    "gameServerStart",
			"time":    t,
			"objtype": "server",
			"obj":     "gameServer",
			"content": "start",
		})
	if err != nil {
		e.log.Warnf("写入数据库异常: %s", err.Error)
	}
}
func (e *even) saveSay(t string, obj string, content string) {
	obj = obj[1 : len(obj)-1]
	nchan.CoverString(server.Server.Status["message"].(chan string), fmt.Sprintf("%s:%s", obj, content))
	err := database.DataBases["even"].Write(
		map[string]string{
			"name":    "say",
			"time":    t,
			"objtype": "users",
			"obj":     obj,
			"content": content,
		})
	if err != nil {
		e.log.Warnf("写入数据库异常: %s", err.Error)
	}
}
func (e *even) saveStop(t string) {
	nchan.CoverString(server.Server.Status["message"].(chan string), "服务器已关闭")
	err := database.DataBases["even"].Write(
		map[string]string{
			"name":    "gameServerStop",
			"time":    t,
			"objtype": "server",
			"obj":     "gameServer",
			"content": "stop",
		})
	if err != nil {
		e.log.Warnf("写入数据库异常: %s", err.Error)
	}
}
func (e *even) saveLogin(t string, obj string) {

	var content string = ""
	info := strings.Split(obj, "[")
	obj = obj[:len(obj)-len(info[len(info)-1])-1]
	content = info[len(info)-1][1 : len(info[len(info)-1])-1]
	server.Server.Status["users"].(map[string]interface{})["online"] =
		server.Server.Status["users"].(map[string]interface{})["online"].(int) + 1
	server.Server.Status["users"].(map[string]interface{})["userlist"] =
		append(server.Server.Status["users"].(map[string]interface{})["userlist"].([]string), obj)
	nchan.CoverString(server.Server.Status["message"].(chan string), fmt.Sprintf("%s 上线了", obj))
	err := database.DataBases["even"].Write(
		map[string]string{
			"name":    "login",
			"time":    t,
			"objtype": "users",
			"obj":     obj,
			"content": content,
		})
	if err != nil {
		e.log.Warnf("写入数据库异常: %s", err.Error)
	}
}
func (e *even) saveLogout(t string, obj string) {
	server.Server.Status["users"].(map[string]interface{})["online"] =
		server.Server.Status["users"].(map[string]interface{})["online"].(int) - 1
	nchan.CoverString(server.Server.Status["message"].(chan string), fmt.Sprintf("%s 下线了", obj))
	err := database.DataBases["even"].Write(
		map[string]string{
			"name":    "logout",
			"time":    t,
			"objtype": "users",
			"obj":     obj,
			"content": "",
		})
	if err != nil {
		e.log.Warnf("写入数据库异常: %s", err.Error)
	}
}