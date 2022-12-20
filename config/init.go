package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Setup() {
	setEnv()
	setGinMode()
}

func PrivateKey() []byte {
	return []byte(os.Getenv("JWT_PRIVATE_KEY"))
}
