package config

import "os"

type Config struct {
	// конфиг микросервиса
	ServerPort  string
	ServiceName string

	// данные для бд
	PGHost   string
	PGPort   string
	PGUser   string
	PGPass   string
	PGDBName string
}

func Load() (*Config, error) {
	return &Config{
		ServerPort:  getEnv("SERVER_PORT", "5050"),
		ServiceName: "User",
		PGHost:      getEnv("PG_HOST", "localhost"),
		PGPort:      getEnv("PG_PORT", "5432"),
		PGUser:      getEnv("PG_USER", "postgres"),
		PGPass:      getEnv("PG_PASSWORD", "postgres"),
		PGDBName:    getEnv("PG_NAME", "user"),
	}, nil
}

func getEnv(key, defaultValue string) string {

	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
