package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	MongoDbUrl  string
	MongoDbName string
)

func Init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MongoDbUrl = os.Getenv("MongoDbUrl")
	MongoDbName = os.Getenv("MongoDbName")
}
