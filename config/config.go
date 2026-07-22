package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBName     string
	DBPort     string
	DBPassword string
	DBUser     string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Failed to Load .env File. Falling Back to System Environment Variables and Default Configs")
	}

	return &Config{
		DBHost:     getEnv("DB_HOST"),
		DBName:     getEnv("DB_NAME"),
		DBPort:     getEnv("DB_PORT"),
		DBPassword: getEnv("DB_PASSWORD"),
		DBUser:     getEnv("DB_USER"),
	}
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatalf("Key: %s Does is not present in environ", key)
	return ""
}
