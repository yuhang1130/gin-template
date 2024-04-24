package middleware

import (
	"gin-template/auth"
	"gin-template/common"
	"gin-template/common/enum"
	"gin-template/common/errorCode"
	"gin-template/global"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func VerifyJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		code := errorCode.SUCCESS
		authorization := c.GetHeader(global.Config.Jwt.Name)
		if authorization == "" {
			global.Log.Warnf("authorization is empty. authorization: %s \n", authorization)
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.Result{Code: code, Msg: "Authorization header is missing"})
			c.Abort()
			return
		}

		// 检查Bearer token格式
		bearerToken := strings.Split(authorization, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			global.Log.Warnf("Invalid Authorization header format. authorization: %s \n", authorization)
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.Result{Code: code, Msg: "Invalid Authorization header format"})
			c.Abort()
			return
		}

		// 解析JWT
		tokenString := bearerToken[1]
		claims, err := auth.ParseToken(tokenString)
		if err != nil {
			code = errorCode.UNKNOW_IDENTITY
			c.JSON(http.StatusUnauthorized, common.Result{Code: code, Msg: err.Error()})
			c.Abort()
			return
		}
		global.Log.Infoln("parse token success sessionID:", claims.SessionID)
		c.Set(enum.SessionID, claims.SessionID)
		c.Next()
	}
}
