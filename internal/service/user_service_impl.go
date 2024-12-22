package service

import (
	"github.com/NeoJay0705/go-server-template/internal/model"
	"github.com/NeoJay0705/go-server-template/internal/repository"
)

type UserServiceImpl struct {
	repo repository.UserTemplate
}

func NewUserService(repo repository.UserTemplate) UserService {
	return &UserServiceImpl{repo: repo}
}

func (u *UserServiceImpl) GetUser(id int64) (*model.User, error) {
	return u.repo.FindUser(id)
}
