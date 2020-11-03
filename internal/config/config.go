package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Host 		string
	Port 		string
	PublicDir   string
	SimfileDir 	string
}

type Config struct {
	Server ServerConfig
}

// New returns a new Config struct
func New() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	}
	return &Config{
		Server: ServerConfig{
			Host: getEnv("HOST", ""),
			Port: getEnv("PORT", ""),
			PublicDir: getEnv("PUBLIC_DIR", ""),
			SimfileDir: getEnv("SIMFILE_DIR", ""),
		},
	}
}

// getEnv helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}