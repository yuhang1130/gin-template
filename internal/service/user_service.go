package service

import (
	"encoding/json"
	"fmt"
	"gin-template/auth"
	"gin-template/common/enum"
	errorCode "gin-template/common/error_code"
	"gin-template/common/utils"
	"gin-template/global"
	"gin-template/internal/api/request"
	"gin-template/internal/api/response"
	"gin-template/internal/dao"
	"gin-template/internal/model"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Register(ctx *gin.Context, dto request.UserCreateDto) error
	Login(ctx *gin.Context, dto request.UserLoginDto) (*response.UserLoginResDto, error)
	Logout(ctx *gin.Context) (bool, error)
	Info(ctx *gin.Context) (*enum.PartialUser, error)
}

type UserImpl struct {
	dao dao.UserRepo
}

func NewUserService(dao dao.UserRepo) UserService {
	return &UserImpl{dao: dao}
}

func (ui *UserImpl) Register(ctx *gin.Context, data request.UserCreateDto) error {
	// 1.查询用户是否存在
	_, err := ui.dao.GetByUserName(ctx, data.UserName)
	if err == nil {
		return errorCode.Error_ACCOUNT_EXISTS
	}

	// 1.新增员工,构建员工基础信息
	entity := model.User{
		UserName: data.UserName,
		Phone:    data.Phone,
		Email:    data.Email,
	}
	// 新增用户为启用状态
	salt, saltErr := utils.GenerateRandomString(16) // 生成一个长度为 16 的随机盐
	if saltErr != nil {
		panic(err)
	}
	global.Log.Errorf("Generated Salt: %s", salt)
	entity.Salt = salt
	entity.Password = utils.MD5V(data.Password, salt, 1)
	// 新增用户
	if insertErr := ui.dao.Insert(ctx, entity); insertErr != nil {
		return insertErr
	}
	return nil
}

func (ui *UserImpl) Login(ctx *gin.Context, data request.UserLoginDto) (*response.UserLoginResDto, error) {
	// 1.查询用户是否存在
	user, err := ui.dao.GetByUserName(ctx, data.UserName)
	if err != nil || user == nil {
		return nil, errorCode.Error_ACCOUNT_NOT_FOUND
	}
	// 2.校验密码
	password := utils.MD5V(data.Password, user.Salt, 1)
	if password != user.Password {
		return nil, errorCode.Error_PASSWORD_ERROR
	}

	// 3.校验状态
	if user.Status == enum.DISABLE {
		return nil, errorCode.Error_ACCOUNT_LOCKED
	}

	// 4.存会话信息
	UserSession := enum.SessionDto{
		UserId:   user.ID,
		OpUserId: user.ID,
		User: enum.PartialUser{
			UserName: user.UserName,
			IsAdmin:  user.IsAdmin,
			Phone:    user.Phone,
			Email:    user.Email,
		},
		Rights: []string{"test", "test1"},
	}
	// 将用户数据转换为JSON字符串
	UserSessionJSON, _ := json.Marshal(UserSession)

	session, _ := global.RedisStore.Get(ctx.Request, enum.Sid)
	// 将JSON数据存储为字符串
	session.Values[enum.SessionData] = string(UserSessionJSON)
	// 保存更改
	if err := session.Save(ctx.Request, ctx.Writer); err != nil {
		global.Log.Errorf("Error saving session: %v", err.Error())
		return &response.UserLoginResDto{}, errorCode.Error_SESSION_SAVE_FAILED
	}

	fmt.Println("session.ID--------", session.ID)
	token, err := auth.GenerateToken(session.ID)
	if err != nil {
		global.Log.Errorf("GenerateToken fail. msg: %s", err.Error())
		return &response.UserLoginResDto{}, errorCode.Error_LOGIN_FAILED
	}

	// 5.构造返回数据
	resp := response.UserLoginResDto{
		Token: token,
	}
	return &resp, nil

}

func (ui *UserImpl) Logout(ctx *gin.Context) (bool, error) {
	sessionID := ctx.MustGet(enum.SessionID)
	sessionData, err := utils.GetSessionData(ctx)
	if err != nil {
		global.Log.Errorf("GetSessionData fail. msg: %s", err.Error())
		return false, err
	}
	userName := sessionData.User.UserName

	// 获取session
	session, _ := global.RedisStore.Get(ctx.Request, enum.Sid)
	// 删除用户信息
	delete(session.Values, enum.SessionData)
	// 保存更改
	if err := session.Save(ctx.Request, ctx.Writer); err != nil {
		global.Log.Errorf("Logout Error saving session. err: %s", err.Error())
		return false, err
	}

	// 清除cookie
	ctx.SetCookie(enum.Sid, "", -1, "/", "", false, true)
	global.Log.Warnf("Logout Success. sessionID: %s, userName: %s", sessionID, userName)
	return true, nil
}

func (ui *UserImpl) Info(ctx *gin.Context) (*enum.PartialUser, error) {
	sessionData, err := utils.GetSessionData(ctx)
	if err != nil {
		global.Log.Errorf("GetSessionData fail. msg: %s", err.Error())
		return &enum.PartialUser{}, err
	}

	return &sessionData.User, nil
}
