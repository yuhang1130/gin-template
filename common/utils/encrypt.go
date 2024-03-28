package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

// GenerateRandomBytes 返回指定长度的随机字节切片
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// 注意：err != nil 时，len(b) == n，但 b 未被初始化
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GenerateRandomString 返回指定长度的随机字符串，使用 URL 安全的 base64 编码
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

// MD5V 对目标字符串取Hash slat：加盐字段，iteration：hash迭代轮数。
func MD5V(str string, salt string, iteration int) string {
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s) // 先传入盐值，之前因为顺序错了卡了很久
	h.Write(b)
	var res []byte
	res = h.Sum(nil)
	for i := 0; i < iteration-1; i++ {
		h.Reset()
		h.Write(res)
		res = h.Sum(nil)
	}
	return hex.EncodeToString(res)
}
