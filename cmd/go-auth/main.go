package main

import (
	"fmt"

	"github.com/enespolat24/go-auth/internal/config"
	"github.com/enespolat24/go-auth/internal/db"
)

func main() {
	config.GetConfig()
	db := db.NewConnection()
	fmt.Println(db)
}
