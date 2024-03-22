package service

import (
	"context"
	"fmt"
	"gin-template/internal/api/request"
	"gin-template/internal/dao"
	"gin-template/internal/model"
)

type UserService interface {
	Register(ctx context.Context, dto request.UserCreateDto) (*model.User, error)
	// login(ctx context.Context, dto request.UserLoginDto) (*response.UserLoginResDto, error)
	// logout(ctx context.Context, user model.User) (bool, error)
	// info(ctx context.Context, user model.User) (*model.User, error)
}

type UserImpl struct {
	dao dao.UserRepo
}

func NewUserService(dao dao.UserRepo) UserService {
	return &UserImpl{dao: dao}
}

func (s *UserImpl) Register(ctx context.Context, dto request.UserCreateDto) (*model.User, error) {

	fmt.Printf("UserImpl------dto----: %+v \n", dto)
	user, err := s.dao.Insert(ctx, dto)
	if err != nil {
		return nil, err
	}
	return user, nil
}
