package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret   string
	FrontendURL string
	Port        string

	PostgresURL string
}

func (config *Config) IsValid() bool {
	return len(config.JWTSecret) != 0 &&
		len(config.FrontendURL) != 0 &&
		len(config.Port) != 0
}

var Vars *Config

func Initialize(devMode bool) {

	if devMode {
		godotenv.Load(".env.local")
		godotenv.Load(".env.development")
	} else {
		godotenv.Load(".env.local")
		godotenv.Load(".env.production")
	}

	godotenv.Load(".env")

	Vars = new(Config)
	Vars.JWTSecret = os.Getenv("JWT_SECRET")
	Vars.FrontendURL = os.Getenv("FRONTEND_URL")
	Vars.Port = os.Getenv("PORT")
	Vars.PostgresURL = os.Getenv("POSTGRES_URL")

	if !Vars.IsValid() {
		panic("failed to initialize config")
	}
}
