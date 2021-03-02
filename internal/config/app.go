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

	// MongoConnectionTimeout is MongoDB connection timeout duration (second)
	MongoConnectionTimeout int

	// MongoMaxPoolSize max connection pool size
	MongoMaxPoolSize uint64
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	MongoURL = getEnvWithDefault("MONGO_URL", "mongodb://127.0.0.1:27017")
	MongoTodoDbName = getEnvWithDefault("MONGO_TODO_DB", "todos")

	if val, err := strconv.Atoi(getEnvWithDefault("MONGO_CONNECTION_TIMEOUT", "10")); err != nil {
		MongoConnectionTimeout = 10
	} else {
		MongoConnectionTimeout = val
	}

	if val, err := strconv.ParseUint(getEnvWithDefault("MONGO_CONNECTION_TIMEOUT", "50"), 10, 64); err != nil {
		MongoMaxPoolSize = 10
	} else {
		MongoMaxPoolSize = val
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
