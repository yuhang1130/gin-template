package auth

import (
	"fmt"
	"gin-template/common/errorCode"
	"gin-template/global"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	SessionID string `json:"sessionID"`
}

// 生成JWT令牌
func GenerateToken(sessionID string) (string, error) {
	expireTime := time.Duration(global.Config.Jwt.TTL) * time.Second
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)),
		},
		SessionID: sessionID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名token
	tokenString, err := token.SignedString([]byte(global.Config.Jwt.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func ParseToken(tokenString string) (*CustomClaims, error) {
	claims := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(global.Config.Jwt.Secret), nil
	})

	if err != nil {
		global.Log.Warnf("jwt parse token fail! msg: %s", err.Error())
		return nil, err
	}
	if !token.Valid {
		global.Log.Warnf("invalid token. token: %s", tokenString)
		return nil, errorCode.Error_INVALID_TOKEN
	}

	return claims, nil
}
