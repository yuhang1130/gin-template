package controller

import (
	"gin-template/common"
	errorCode "gin-template/common/error_code"
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

func (uc *UserController) Register(ctx *gin.Context) {
	code := errorCode.SUCCESS
	userDto := request.UserCreateDto{}
	if err := ctx.Bind(&userDto); err != nil {
		code = errorCode.ERROR
		global.Log.Error("UserController Register 解析失败")
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}

	err := uc.service.Register(ctx, userDto)
	if err != nil {
		code = errorCode.ERROR
		global.Log.Warnf("UserController Register Error: %s \n", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, common.Result{
		Msg:  "success",
		Data: true,
		Code: code,
	})
}

func (uc *UserController) Login(ctx *gin.Context) {
	code := errorCode.SUCCESS
	userLogin := request.UserLoginDto{}
	err := ctx.Bind(&userLogin)
	if err != nil {
		code = errorCode.ERROR
		global.Log.Warnf("UserController Login 解析失败 Error: %s \n", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	resp, err := uc.service.Login(ctx, userLogin)
	if err != nil {
		code = errorCode.ERROR
		global.Log.Warnf("UserController Login Error: %s \n", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: resp,
	})
}

func (uc *UserController) Logout(ctx *gin.Context) {
	code := errorCode.SUCCESS
	resp, err := uc.service.Logout(ctx)
	if err != nil {
		code = errorCode.ERROR
		global.Log.Warnf("UserController Logout Error: %s \n", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: resp,
	})
}

func (uc *UserController) Info(ctx *gin.Context) {
	code := errorCode.SUCCESS
	resp, err := uc.service.Info(ctx)
	if err != nil {
		code = errorCode.ERROR
		global.Log.Warnf("UserController Info Error: %s \n", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: resp,
	})
}
