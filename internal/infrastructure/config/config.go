package config

import (
	"log"
	"os"
)

type Config struct {
	APPEnv      string
	AppPort     string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_Password string
	DB_Name     string
	DB_SSLMode  string
	JWT_Secret  string
}

func LoadConfig() *Config {
	return &Config{
		APPEnv:      getEnv("APP_ENV", "development"),
		AppPort:     getEnv("APP_PORT", "8000"),
		DB_HOST:     getEnv("DB_HOST", "localhost"),
		DB_PORT:     getEnv("DB_PORT", "5432"),
		DB_USER:     getEnv("DB_USER", "postgres"),
		DB_Password: getEnv("DB_PASSWORD", ""),
		DB_Name:     getEnv("DB_NAME", "hospital_db"),
		DB_SSLMode:  getEnv("DB_SSLMODE", "disable"),
		JWT_Secret:  getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, value string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Print("env variable not found ", key)
		return value
	}
	return val
}
