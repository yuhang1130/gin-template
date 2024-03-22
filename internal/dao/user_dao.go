package dao

import (
	"context"
	"fmt"
	"gin-template/internal/api/request"
	"gin-template/internal/model"

	"gorm.io/gorm"
)

type UserRepo interface {
	Insert(ctx context.Context, dto request.UserCreateDto) (*model.User, error)
	// login(ctx context.Context, dto request.UserLoginDto) (response.UserLoginResDto, error)
	// logout(ctx context.Context, user model.User) (bool, error)
	// info(ctx context.Context, user model.User) (*model.User, error)
}

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserRepo {
	return &UserDao{db: db}
}

func (c *UserDao) Insert(ctx context.Context, dto request.UserCreateDto) (*model.User, error) {
	fmt.Printf("UserDao------Insert-----: %+v \n", dto)
	return &model.User{}, nil
}

// func (c *UserDao) login(ctx context.Context, id uint64) (response.UserLoginResDto, error) {
// }

// func (c *UserDao) logout(ctx context.Context, id uint64) (bool, error) {
// }

// func (c *UserDao) info(ctx context.Context, id uint64) (model.User, error) {
// }
