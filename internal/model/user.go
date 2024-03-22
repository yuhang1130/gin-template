package model

import (
	"gin-template/common/enum"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string `gorm:"type:varchar(50);not null;unique;comment:'用户名'" json:"userName"`
	Password   string `gorm:"size:255;not null;comment:'用户密码'" json:"passWord"`
	Salt       string `gorm:"type:varchar(255);comment:'盐'" json:"salt"`
	IsAdmin    string `gorm:"type:varchar(100);comment:'是否为管理员'" json:"isAdmin"`
	Phone      string `gorm:"type:varchar(15);not null;comment:'手机号'" json:"phone"`
	Email      string `gorm:"type:varchar(100);comment:'邮箱'" json:"email"`
	CreateUser uint64 `gorm:"type:uint(64);comment:'创建者'" json:"createUser"`
	UpdateUser uint64 `gorm:"type:uint(64);comment:'更新者'" json:"updateUser"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	// 自动填充 创建时间、创建人、更新时间、更新用户
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	// 从上下文获取用户信息
	value := tx.Statement.Context.Value(enum.CurrentId)
	if uid, ok := value.(uint64); ok {
		u.CreateUser = uid
		u.UpdateUser = uid
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	// 在更新记录千自动填充更新时间
	u.UpdatedAt = time.Now()
	// 从上下文获取用户信息
	value := tx.Statement.Context.Value(enum.CurrentId)
	if uid, ok := value.(uint64); ok {
		u.UpdateUser = uid
	}
	return nil
}
