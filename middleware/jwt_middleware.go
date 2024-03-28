package middleware

import (
	"gin-template/auth"
	"gin-template/common"
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
			code = errorCode.UNKNOW_IDENTITY
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.Result{Code: code, Msg: "Authorization header is missing"})
			return
		}

		// 检查Bearer token格式
		bearerToken := strings.Split(authorization, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			code = errorCode.UNKNOW_IDENTITY
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.Result{Code: code, Msg: "Invalid Authorization header format"})
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
		c.Set("sessionID", claims["sessionID"])
		c.Next()
	}
}
