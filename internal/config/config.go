package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// Database holds configuration for database.
type Database struct {
	Host         string `env:"DATABASE_HOST,default=localhost"`
	Port         string `env:"DATABASE_PORT,default=5432"`
	Username     string `env:"DATABASE_USERNAME,required"`
	Password     string `env:"DATABASE_PASSWORD,required"`
	Name         string `env:"DATABASE_NAME,required"`
	SSLMode      string `env:"DATABASE_SSL_MODE,default=disable"`
	MaxOpenConns int    `env:"DATABASE_MAX_OPEN_CONNS,default=5"`
	MaxIdleConns int    `env:"DATABASE_MAX_IDLE_CONNS,default=1"`
}

// Google holds configuration related to Google.
type Google struct {
	Audience string `env:"GOOGLE_AUDIENCE,required"`
}

// Config holds configuration for the project.
type Config struct {
	Port     string `env:"PORT,default=6666"`
	Database Database
	Google   Google
}

// NewConfig creates an instance of Config.
// It needs the path of the env file to be used.
func NewConfig(env string) (*Config, error) {
	godotenv.Load(env)

	var config Config
	if err := envdecode.Decode(&config); err != nil {
		return nil, errors.Wrap(err, "[NewConfig] error decoding env")
	}

	return &config, nil
}
