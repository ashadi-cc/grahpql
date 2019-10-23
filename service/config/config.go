package config

import (
	"fmt"
	"log"
	"time"

	"github.com/caarlos0/env"
)

//Config configuration model
type Config struct {
	AppName     string        `env:"APP_NAME" envDefault:"grahpql"`
	AppPort     string        `env:"APP_PORT" envDefault:"8001"`
	DBHost      string        `env:"DB_HOST" envDefault:"localhost"`
	DBPort      string        `env:"DB_PORT" envDefault:"5432"`
	DBUser      string        `env:"DB_USER" envDefault:"postgres"`
	DBPassword  string        `env:"DB_PASSWORD" envDefault:"postgres"`
	DBName      string        `env:"DB_NAME" envDefault:"grahpql"`
	JWTSecret   string        `env:"JWT_SECRET" envDefault:"secret"`
	JWTExpireIn time.Duration `env:"JWT_EXPIRED_IN" envDefault:"40000s"`
	DebugMode   bool          `env:"DEBUG_MODE" envDefault:"true"`
	LogFormat   string        `env:"LOG_FORMAT" envDefault:"%{color}%{time:2006/01/02 15:04:05 -07:00 MST} [%{level:.6s}] %{shortfile} : %{color:reset}%{message}"`
}

var config *Config

//GetConfig getConfig variable
func GetConfig() *Config {
	if config == nil {
		config = &Config{}
		err := env.Parse(config)
		if err != nil {
			log.Fatal("error when parse configuration ", err.Error())
		}
		config.AppPort = fmt.Sprintf(":%s", config.AppPort)
	}

	return config
}
