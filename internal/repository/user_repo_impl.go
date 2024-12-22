package repository

import (
	"github.com/NeoJay0705/go-server-template/internal/model"
	"gorm.io/gorm"
)

type UserTemplateImpl struct {
	db *gorm.DB
}

func NewUserTemplate(db *gorm.DB) UserTemplate {
	return &UserTemplateImpl{db: db}
}

func (templ *UserTemplateImpl) CreateUser(user *model.User) error {
	return templ.db.Create(user).Error
}

func (templ *UserTemplateImpl) FindUser(id int64) (*model.User, error) {
	user := new(model.User)
	user.ID = id
	err := templ.db.First(&user).Error
	return user, err
}
