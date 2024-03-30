package middleware

import (
	"fmt"
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
		fmt.Printf("c-----111-----: %+v \n", c.Request.Cookies())

		// 从请求头中获取token
		code := errorCode.SUCCESS
		authorization := c.GetHeader(global.Config.Jwt.Name)
		if authorization == "" {
			code = errorCode.UNKNOW_IDENTITY
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.Result{Code: code, Msg: "Authorization header is missing"})
			c.Abort()
			return
		}

		// 检查Bearer token格式
		bearerToken := strings.Split(authorization, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			code = errorCode.UNKNOW_IDENTITY
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
		global.Log.Infoln("parseToken sessionID:", claims["sessionID"])
		c.Set("sessionID", claims["sessionID"])
		// 有传token的以token为准，删除cookie中存的sessionID
		c.Next()
	}
}
