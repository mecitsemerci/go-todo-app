package config

//RedisConfiguration stores mongodb connection info
type RedisConfiguration struct {
	// MongoURL is MongoDB connection string
	RedisURL string `env:"REDIS_URL" envDefault:"127.0.0.1:6379"`

	// MongoTodoDbName is MongoDB database name
	RedisDb int `env:"REDIS_TODO_DB" envDefault:"0"`

	// MongoConnectionTimeout is MongoDB connection timeout duration (second)
	RedisConnectionTimeout int `env:"REDIS_CONNECTION_TIMEOUT" envDefault:"20"`

	// MongoMaxPoolSize max connection pool size
	RedisMaxPoolSize int `env:"REDIS_MAX_POOL_SIZE" envDefault:"10"`
}
