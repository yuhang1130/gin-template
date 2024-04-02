package initialize

import (
	"gin-template/global"
	"gin-template/internal/router"
	"gin-template/middleware"

	"github.com/gin-gonic/gin"
)

func routerInit() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	// 使用自定义的 logrus 中间件
	r.Use(middleware.LoggerMiddleware(global.Log))

	allRouter := router.AllRouter

	// admin
	admin := r.Group("/admin")
	{
		allRouter.UserRouter.InitApiRouter(admin)

	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
