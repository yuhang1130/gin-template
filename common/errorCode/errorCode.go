package errorCode

import "errors"

const (
	SUCCESS         = 0   // ok
	ERROR           = 1   // 内部错误
	UNKNOW_IDENTITY = 401 // 未知身份
)

var (
	Error_PASSWORD_ERROR       = errors.New("密码错误")
	Error_ACCOUNT_NOT_FOUND    = errors.New("账号或密码错误")
	Error_ACCOUNT_LOCKED       = errors.New("账号被锁定")
	Error_ACCOUNT_EXISTS       = errors.New("账号已存在")
	Error_UNKNOWN_ERROR        = errors.New("未知错误")
	Error_USER_NOT_LOGIN       = errors.New("用户未登录")
	Error_LOGIN_FAILED         = errors.New("登录失败")
	Error_UPLOAD_FAILED        = errors.New("文件上传失败")
	Error_PASSWORD_EDIT_FAILED = errors.New("密码修改失败")
	Error_SESSION_SAVE_FAILED  = errors.New("session会话存储失败")
	Error_TOKEN_EXPIRE         = errors.New("token已过期")
	Error_INVALID_TOKEN        = errors.New("token无效")
)
