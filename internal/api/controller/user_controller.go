package controller

import (
	"fmt"
	"gin-template/common"
	"gin-template/common/errorCode"
	"gin-template/global"
	"gin-template/internal/api/request"
	"gin-template/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) Register(ctx *gin.Context) {
	code := errorCode.SUCCESS
	var userDto request.UserCreateDto
	err := ctx.Bind(&userDto)
	if err != nil {
		global.Log.Errorf("param UserCreateDto json failed: %+v \n", err.Error())
		return
	}
	fmt.Printf("userDto----UserController----: %+v \n", userDto)
	if _, err = c.service.Register(ctx, userDto); err != nil {
		code = errorCode.ERROR
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
	})
}
