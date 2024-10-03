package config

type App struct {
	Name           string `env:"APP_NAME" envDefault:"Todo App"`
	Version        string `env:"APP_VERSION" envDefault:"1.0.0"`
	Port           int    `env:"APP_PORT" envDefault:"8080"`
	Host           string `env:"APP_HOST" envDefault:"localhost"`
	JWTSecret      string `env:"JWT_SECRET" envDefault:"secret"`
	JWTExpiresSec  int    `env:"JWT_EXPIRES_HOUR" envDefault:"3600"`
}
