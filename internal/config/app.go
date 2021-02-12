package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	MongoDbUrl        string
	MongoDbTodoDbName string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	MongoDbUrl = getEnvWithDefault("MONGODB_URL", "mongodb://127.0.0.1:27017")
	MongoDbTodoDbName = getEnvWithDefault("MONGODB_TODO_DB", "todos")
}

func getEnvWithDefault(key string, defaultValue string) string {
	var env string
	if value, ok := os.LookupEnv(key); !ok {
		env = defaultValue
	} else {
		env = value
	}
	return env
}
