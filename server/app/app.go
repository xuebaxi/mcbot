package app

import (
	"github.com/gin-gonic/gin"
)

//RegisterRoutes 用于注册路由
func RegisterRoutes(g *gin.Engine) {
	new(com).RegisterRoutes(g)
	new(message).RegisterRoutes(g)
}
