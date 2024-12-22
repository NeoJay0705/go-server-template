package service

import (
	"github.com/NeoJay0705/go-server-template/internal/model"
)

type UserService interface {
	GetUser(id int64) (*model.User, error)
}
