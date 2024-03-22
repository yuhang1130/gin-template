package initialize

import (
	"gin-template/config"
	"gin-template/global"
	"gin-template/logger"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	// 配置文件初始化
	global.Config = config.InitLoadConfig()
	// Log初始化
	global.Log = logger.InitLogger(global.Config.Log.Level)
	// Gorm初始化
	global.DB = InitDatabase()
	// Redis初始化
	global.Redis = initRedis()
	// Router初始化
	router := routerInit()
	return router
}
