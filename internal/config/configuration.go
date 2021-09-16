package config

import (
	"io/ioutil"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	//MongoConfig provides mongodb connection info
	MongoConfig MongoConfiguration

	//RedisConfig provides mongodb connection info
	RedisConfig RedisConfiguration
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	// log only debug mode
	if gin.Mode() != gin.DebugMode {
		log.SetOutput(ioutil.Discard)
	}
}

//Load initialize environment variables
func Load() {
	_ = env.Parse(&RedisConfig)
	_ = env.Parse(&MongoConfig)
}
