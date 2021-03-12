package config

import "github.com/caarlos0/env/v6"

//MongoConfig provides mongodb connection info
var MongoConfig MongoConfiguration

//Load initialize environment variables
func Load() error {
	return env.Parse(&MongoConfig)
}
