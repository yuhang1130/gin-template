package admin

import (
	"gin-template/global"
	"gin-template/internal/api/controller"
	"gin-template/internal/dao"
	"gin-template/internal/service"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (u *UserRouter) InitApiRouter(router *gin.RouterGroup) {
	publicRouter := router.Group("user")
	// privateRouter := router.Group("user")
	// 私有路由使用jwt验证
	// privateRouter.Use(middle.VerifyJWTAdmin())
	// 依赖注入
	userRepo := dao.NewUserDao(global.DB)
	employeeSvc := service.NewUserService(userRepo)
	employeeCtl := controller.NewUserController(employeeSvc)
	{
		publicRouter.POST("/register", employeeCtl.Register)
		// publicRouter.POST("/login", employeeCtl.Login)
		// privateRouter.POST("/logout", employeeCtl.Logout)
		// privateRouter.POST("", employeeCtl.AddEmployee)
		// privateRouter.POST("/status/:status", employeeCtl.OnOrOff)
		// privateRouter.PUT("/editPassword", employeeCtl.EditPassword)
		// privateRouter.PUT("", employeeCtl.UpdateEmployee)
		// privateRouter.GET("/page", employeeCtl.PageQuery)
		// privateRouter.GET("/:id", employeeCtl.GetById)
	}
}
