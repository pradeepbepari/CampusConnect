package config

import (
	"embed"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost            string `env:"DB_HOST"`
	DBPort            string `env:"DB_PORT"`
	DBUser            string `env:"DB_USER"`
	DBPassword        string `env:"DB_PASSWORD"`
	DBName            string `env:"DB_NAME"`
	ServerPort        string `env:"SERVER_PORT"`
	MySQLDatabase     string `env:"MYSQL_DATABASE"`
	MySQLRootPassword string `env:"MYSQL_ROOT_PASSWORD"`
}

func LoadConfig(configs embed.FS) (*Config, error) {
	envFile, err := configs.ReadFile("configs/configs.env")
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded env file: %w", err)
	}

	envMap, err := godotenv.Unmarshal(string(envFile))
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal env file: %w", err)
	}

	for k, v := range envMap {
		os.Setenv(k, v)
	}

	config := &Config{
		DBHost:            os.Getenv("DB_HOST"),
		DBPort:            os.Getenv("DB_PORT"),
		DBUser:            os.Getenv("DB_USER"),
		DBPassword:        os.Getenv("DB_PASSWORD"),
		DBName:            os.Getenv("DB_NAME"),
		ServerPort:        os.Getenv("SERVER_PORT"),
		MySQLDatabase:     os.Getenv("MYSQL_DATABASE"),
		MySQLRootPassword: os.Getenv("MYSQL_ROOT_PASSWORD"),
	}
	return config, nil
}
