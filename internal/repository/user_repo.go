package repository

import (
	"github.com/NeoJay0705/go-server-template/internal/model"
)

type UserTemplate interface {
	CreateUser(user *model.User) error
	FindUser(id int64) (*model.User, error)
}
