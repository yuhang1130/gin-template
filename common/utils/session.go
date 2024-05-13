package utils

import (
	"encoding/json"
	"gin-template/common/enum"
	errorCode "gin-template/common/error_code"
	"gin-template/global"

	"github.com/gin-gonic/gin"
)

func GetSessionData(ctx *gin.Context) (*enum.SessionDto, error) {
	var sessionData = &enum.SessionDto{}
	// 获取session
	session, _ := global.RedisStore.Get(ctx.Request, enum.Sid)
	sessionDataJSON, ok := session.Values[enum.SessionData].(string)
	if !ok {
		global.Log.Error("sessionDataJSON字符串获取失败！")
		return sessionData, errorCode.Error_TOKEN_EXPIRE
	}

	// 将JSON字符串解码回原始的数据结构
	err := json.Unmarshal([]byte(sessionDataJSON), &sessionData)
	if err != nil {
		global.Log.Error("sessionDataJSON字符串解码失败！")
		return sessionData, errorCode.Error_TOKEN_EXPIRE
	}

	return sessionData, nil
}
