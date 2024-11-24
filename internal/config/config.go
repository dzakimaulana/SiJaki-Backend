package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadDBConfig() *DBConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("💩 Error loading .env file")
	}

	return &DBConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
}

type JWTConfig struct {
	Secret string
}

func LoadJWTConfig() *JWTConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("💩 Error loading .env file")
	}

	return &JWTConfig{
		Secret: os.Getenv("JWT_SECRET"),
	}
}

type AppConfig struct {
	Port     string
	LogLevel string
}

func LoadAppConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("💩 Error loading .env file")
	}

	return &AppConfig{
		Port:     os.Getenv("APP_PORT"),
		LogLevel: os.Getenv("LOG_LEVEL"),
	}
}

type MqttConfig struct {
	Broker   string
	Port     string
	Username string
	Password string
}

func LoadMqttConfig() *MqttConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("💩 Error loading .env file")
	}

	return &MqttConfig{
		Broker:   os.Getenv("MQTT_BROKER"),
		Port:     os.Getenv("MQTT_PORT"),
		Username: os.Getenv("MQTT_USERNAME"),
		Password: os.Getenv("MQTT_PASSWORD"),
	}
}
