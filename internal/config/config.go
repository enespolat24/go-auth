package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func GetConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	fmt.Println("env loaded successfully")
}
