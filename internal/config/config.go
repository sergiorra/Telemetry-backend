package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config representation of config into struct
type Config struct {
	Server ServerConfig
}

// ServerConfig representation of serverConfig into struct
type ServerConfig struct {
	Host 		string
	Port 		string
	PublicDir   string
	SimfileDir 	string
}

// New returns a new Config struct
func New() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	}
	return &Config{
		Server: ServerConfig{
			Host: getEnv("HOST", "0.0.0.0"),
			Port: getEnv("PORT", "3000"),
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