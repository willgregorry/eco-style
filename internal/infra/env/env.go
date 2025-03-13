package env

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"time"
)

type Env struct {
	AppPort string `env:"APP_PORT"`

	DBUsername string `env:"DB_USERNAME"`
	DBPassword string `env:"DB_PASSWORD"`
	DBHost     string `env:"DB_HOST"`
	DBName     string `env:"DB_NAME"`
	DBPort     string `env:"DB_PORT"`

	JWTSecret  string        `env:"JWT_SECRET"`
	JWTExpired time.Duration `env:"JWT_EXPIRED"`
}

func New() (*Env, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	_env := new(Env)
	err = env.Parse(_env)
	if err != nil {
		return nil, err
	}

	return _env, nil
}
