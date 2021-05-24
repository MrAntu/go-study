package routers

import (
	"bluebell/controllers"
	"bluebell/logger"
	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	// 使用中间层，拦截gin log
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册
	r.POST("/signup", controllers.SignUpHandler)
	// 登录
	r.POST("/login", controllers.LoginHandler)
	return r
}
