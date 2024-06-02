package config

//MongoConfiguration stores mongodb connection info
type MongoConfiguration struct {
	// MongoURL is MongoDB connection string
	MongoURL string `env:"MONGO_URL" envDefault:"mongodb://127.0.0.1:27017"`

	// MongoTodoDbName is MongoDB database name
	MongoTodoDbName string `env:"MONGO_TODO_DB" envDefault:"TodoDB"`

	// MongoConnectionTimeout is MongoDB connection timeout duration (second)
	MongoConnectionTimeout int `env:"MONGO_CONNECTION_TIMEOUT" envDefault:"20"`

	// MongoMaxPoolSize max connection pool size
	MongoMaxPoolSize uint64 `env:"MONGO_MAX_POOL_SIZE" envDefault:"10"`
}
