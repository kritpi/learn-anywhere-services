package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	// Postgres
	DB_HOST      string `mapstructure:"DB_HOST"`
	DB_PORT      string `mapstructure:"DB_PORT"`
	DB_NAME      string `mapstructure:"DB_NAME"`
	DB_USER      string `mapstructure:"DB_USER"`
	DB_PASSWORD  string `mapstructure:"DB_PASSWORD"`
	DATABASE_URL string `mapstructure:"DATABASE_URL"`

	// MinIO S3
	MINIO_ROOT_USER     string `mapstructure:"MINIO_ROOT_USER"`
	MINIO_ROOT_PASSWORD string `mapstructure:"MINIO_ROOT_PASSWORD"`
	MINIO_ACCESS_KEY    string `mapstructure:"MINIO_ACCESS_KEY"`
	MINIO_SECRET_KEY    string `mapstructure:"MINIO_SECRET_KEY"`
	MINIO_ENDPOINT      string `mapstructure:"MINIO_ENDPOINT"`
	MINIO_BUCKET        string `mapstructure:"MINIO_BUCKET"`

	// MongoDB
	MONGO_PORT     string `mapstructure:"MONGO_PORT"`
	MONGO_USER     string `mapstructure:"MONGO_USER"`
	MONGO_PASSWORD string `mapstructure:"MONGO_PASSWORD"`
	MONGO_DB       string `mapstructure:"MONGO_DB"`
	MONGO_URL      string `mapstructure:"MONGO_URL"`
}

func NewConfig() *Config {
	config := &Config{}

	viper.SetConfigFile(".env")
	viper.SetConfigType("env") // Explicitly set the type
	viper.AutomaticEnv()       // Enable reading from system environment variables

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("⛔️ Error reading config file:", err)
	}
	fmt.Println("✅ Reading Config")

	if err := viper.Unmarshal(config); err != nil {
		log.Fatalln("⛔️ Unable to decode into struct:", err)
	}
	fmt.Println("✅ Decoded Config")

	return config
}
