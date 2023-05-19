package db

import (
	"os"

	"github.com/enespolat24/go-auth/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func NewConnection() handler {
	dsn := "host=" + os.Getenv("host") + " user=" + os.Getenv("user") + " dbname=" + os.Getenv("dbname") + " port=" + os.Getenv("port") + " sslmode=" + os.Getenv("sslmode") + " TimeZone=" + os.Getenv("timezone")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.User{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	return handler{DB: db}
}
