package main

import (
	"fmt"
	"gin-template/global"
	"gin-template/initialize"

	"github.com/gin-gonic/gin"
)

func main() {
	defer global.RedisStore.Close() // 确保在程序结束时关闭store连接
	router := initialize.Init()
	mode := global.Config.Server.Level
	gin.SetMode(mode)
	port := global.Config.Server.Port
	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		global.Log.Error("Server Run Fail!")
	}
	global.Log.Warn("Server Run Success!")
}
