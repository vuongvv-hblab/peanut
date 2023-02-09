package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const projectDirName = "peanut"

func init() {
	//projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	//currentWorkDirectory, _ := os.Getwd()
	//rootPath := projectName.Find([]byte(currentWorkDirectory))
	//fmt.Println(rootPath)
	//err := godotenv.Load(string(rootPath) + `/.env`)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	getConfig()
}

func Setup() {
	setEnv()
	setGinMode()
}

func PrivateKey() []byte {
	return []byte(os.Getenv("JWT_PRIVATE_KEY"))
}
