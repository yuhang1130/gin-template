package auth

import (
	"fmt"
	"gin-template/common/errorCode"
	"gin-template/global"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// 生成JWT令牌
func GenerateToken(sessionID string) (string, error) {
	// 创建token
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置claims（有效载荷）
	claims := token.Claims.(jwt.MapClaims)
	claims["sessionID"] = sessionID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// 签名token
	tokenString, err := token.SignedString([]byte(global.Config.Jwt.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			// 确保token的方法为JWT签名方法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(global.Config.Jwt.Secret), nil
		},
	)
	if err != nil {
		global.Log.Errorf("jwt.Parse token fail! msg: %s", err.Error())
		return nil, errorCode.Error_PARSE_TOKEN_FAILED
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		global.Log.Error("token.Valid fail!")
		return nil, errorCode.Error_PARSE_TOKEN_FAILED
	}
	// 将用户信息设置到上下文中
	return claims, nil
}
