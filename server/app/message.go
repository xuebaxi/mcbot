package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuebaxi/mcbot/server/command"
)

type message struct {
}

func (m *message) RegisterRoutes(g *gin.Engine) {
	g.GET("/message", m.message)
}

//message 返回系统消息
func (m *message) message(c *gin.Context) {
	c.String(http.StatusOK, command.Command["message"]([]string{}))
}
