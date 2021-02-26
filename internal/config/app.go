package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// MongoURL is MongoDB connection string
	MongoURL string

	// MongoTodoDbName is MongoDB database name
	MongoTodoDbName string

	// DBTimeout is MongoDB connection timeout duration (second)
	DBTimeout int
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	MongoURL = getEnvWithDefault("MONGO_URL", "mongodb://127.0.0.1:27017")
	MongoTodoDbName = getEnvWithDefault("MONGO_TODO_DB", "todos")
	if t, err := strconv.Atoi(getEnvWithDefault("DB_TIMEOUT", "10")); err != nil {
		DBTimeout = t
	}
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
