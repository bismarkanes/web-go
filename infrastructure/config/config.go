package config

import (
    "os"
)

type Config struct {
    AppPort    string
    DBHost     string
    DBUsername string
    DBPassword string
    DBName     string
    DBPort     string
}

func NewConfig() Config {
    config := Config{
        AppPort:    os.Getenv("APP_PORT"),
        DBHost:     os.Getenv("DB_HOST"),
        DBUsername: os.Getenv("DB_USERNAME"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBName:     os.Getenv("DB_NAME"),
        DBPort:     os.Getenv("DB_PORT"),
    }

    if config.AppPort == "" {
        config.AppPort = "8080"
    }

    return config
}
