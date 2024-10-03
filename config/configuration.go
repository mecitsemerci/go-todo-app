package config

import (
	"github.com/caarlos0/env/v11"
)

var (
	AppConfig App
)

// Load initialize environment variables
func init() {
	err := env.Parse(&AppConfig)

	if err != nil {
		panic(err)
	}
}
