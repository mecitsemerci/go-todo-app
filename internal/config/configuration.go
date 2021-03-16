package config

import (
	"github.com/caarlos0/env/v6"
)

var (
	//MongoConfig provides mongodb connection info
	MongoConfig MongoConfiguration

	//RedisConfig provides mongodb connection info
	RedisConfig RedisConfiguration
)

//Load initialize environment variables
func Load() {
	_ = env.Parse(&RedisConfig)
	_ = env.Parse(&MongoConfig)
}
