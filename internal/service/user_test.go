package service

import (
	"testing"

	"github.com/NeoJay0705/go-server-template/internal/repository"
	"github.com/NeoJay0705/go-server-template/pkg/db"
)

func TestGetUser(t *testing.T) {
	usrservice := NewUserService(repository.NewUserTemplate(db.Conn.Pgdb))
	u, err := usrservice.GetUser(1)
	if err != nil {
		t.Error(err)
	}
	t.Log(u)
}
