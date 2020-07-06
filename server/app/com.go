package app

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xuebaxi/mcbot/server/command"
)

type com struct {
}

func (c *com) RegisterRoutes(g *gin.Engine) {
	g.POST("/", c.run)
	g.POST("/admin", c.admin)
}

func (c *com) run(g *gin.Context) {
	content := g.PostForm("input")
	com := strings.Split(content, " ")
	result := string("")
	if function, ok := command.Command[com[0]]; ok {
		result = function(com)
	} else {
		result = "命令没找到"
	}
	g.String(http.StatusOK, result)
}

func (c *com) admin(g *gin.Context) {
	content := g.PostForm("input")
	com := strings.Split(content, " ")
	result := string("")
	if function, ok := command.Admin[com[0]]; ok {
		result = function(com)
	} else {
		result = "命令没找到"
	}
	g.String(http.StatusOK, result)
}
