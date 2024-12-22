package repository

import (
	"testing"

	"github.com/NeoJay0705/go-server-template/pkg/db"
)

func TestInsert(t *testing.T) {
	db := NewUserTemplate(db.Conn.Pgdb)
	user, error := db.FindUser(1)
	if error != nil {
		t.Error(error)
	}
	t.Log(user)
}
