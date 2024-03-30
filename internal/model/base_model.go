package model

import (
	"gin-template/common/enum"

	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	ID         uint              `gorm:"primaryKey;autoIncrement;type:uint"`
	CreateUser uint              `gorm:"type:uint;comment:创建者" json:"createUser"`
	UpdateUser uint              `gorm:"type:uint;comment:更新者" json:"updateUser"`
	DeleteUser uint              `gorm:"type:uint;comment:删除者" json:"deleteUser"`
	Status     enum.CommonStatus `gorm:"type:int;default:1;comment:状态" json:"status"`
}
