package config

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

// Config holds configuration for the project.
type Config struct {
	Port     string `env:"PORT,default=6666"`
	Database Database
}
