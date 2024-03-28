package admin

import (
	"gin-template/global"
	"gin-template/internal/api/controller"
	"gin-template/internal/dao"
	"gin-template/internal/service"
	"gin-template/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (u *UserRouter) InitApiRouter(router *gin.RouterGroup) {
	publicRouter := router.Group("user")
	privateRouter := router.Group("user")
	// 私有路由使用jwt验证
	privateRouter.Use(middleware.VerifyJWT())
	// 依赖注入
	userRepo := dao.NewUserDao(global.DB)
	userSvc := service.NewUserService(userRepo)
	userCtl := controller.NewUserController(userSvc)
	{
		publicRouter.POST("/register", userCtl.Register)
		publicRouter.POST("/login", userCtl.Login)
		privateRouter.POST("/logout", userCtl.Logout)
		// privateRouter.POST("", employeeCtl.AddEmployee)
		// privateRouter.POST("/status/:status", employeeCtl.OnOrOff)
		// privateRouter.PUT("/editPassword", employeeCtl.EditPassword)
		// privateRouter.PUT("", employeeCtl.UpdateEmployee)
		// privateRouter.GET("/page", employeeCtl.PageQuery)
		// privateRouter.GET("/:id", employeeCtl.GetById)
	}
}
