package dao

import (
	"gin-template/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo interface {
	Insert(ctx *gin.Context, entity model.User) error
	GetByUserName(ctx *gin.Context, userName string) (*model.User, error)
	// logout(ctx *gin.Context, user model.User) (bool, error)
	// info(ctx *gin.Context, user model.User) (*model.User, error)
}

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserRepo {
	return &UserDao{db: db}
}

func (u *UserDao) Insert(ctx *gin.Context, entity model.User) error {
	return u.db.WithContext(ctx).Create(&entity).Error
}

func (u *UserDao) GetByUserName(ctx *gin.Context, userName string) (*model.User, error) {
	var user model.User
	err := u.db.WithContext(ctx).Where("user_name=?", userName).First(&user).Error
	return &user, err
}

// func (c *UserDao) login(ctx *gin.Context, id uint64) (response.UserLoginResDto, error) {
// }

// func (c *UserDao) logout(ctx *gin.Context, id uint64) (bool, error) {
// }

// func (c *UserDao) info(ctx *gin.Context, id uint64) (model.User, error) {
// }
