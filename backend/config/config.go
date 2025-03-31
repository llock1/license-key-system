package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret string
}

func (config *Config) IsValid() bool {
	return len(config.JWTSecret) != 0
}

var Vars *Config

func InitConfig(devMode bool) bool {

	if devMode {
		godotenv.Load(".env.development")
	} else {
		godotenv.Load(".env.local")
		godotenv.Load(".env.production")
	}

	godotenv.Load(".env")

	Vars := new(Config)
	Vars.JWTSecret = os.Getenv("JWT_SECRET")

	if !Vars.IsValid() {
		return false
	}

	return true
}
