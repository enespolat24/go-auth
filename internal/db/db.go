package db

import (
	"github.com/enespolat24/go-auth/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func NewConnection(host string, user string, password string, dbName string, port string, sslmode string, timeZone string) handler {
	dsn := "host=" + host + "user=" + user + "dbname=" + dbName + "port=" + port + "sslmode=" + sslmode + "TimeZone=" + timeZone
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.User{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	return handler{DB: db}
}
