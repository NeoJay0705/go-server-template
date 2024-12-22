package db

import (
	"fmt"

	"github.com/NeoJay0705/go-server-template/pkg/config"
	_ "github.com/lib/pq" // PostgreSQL driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	USERNAME = "postgres"
	PASSWORD = "mysecretpassword"
	SERVER   = "127.0.0.1"
	PORT     = 5432
	DATABASE = "postgres"
)

type Config struct {
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"PASSWORD"`
	Server   string `mapstructure:"SERVER"`
	Port     int    `mapstructure:"PORT"`
	Database string `mapstructure:"DATABASE"`
}

type Connection struct {
	Pgdb *gorm.DB
}

var Conn Connection

func init() {
	c := new(Config)
	config.SetConfig(c)
	Conn.Pgdb = connect(c)
}

func connect(config *Config) *gorm.DB {
	// dsn: data source name
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		config.Username, config.Password, config.Server, config.Port, config.Database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func Disconnect() {

}
