package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBUser      string
	DBPass      string
	DBName      string
	DBHost      string // ローカル: "127.0.0.1"; 本番: インスタンス接続名 (PROJECT:REGION:INSTANCE)
	DBPort      string // ローカル環境用
	UseCloudSQL bool
}

func NewConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env ファイルが見つかりませんでした")
	}
	useCloudSQL := os.Getenv("USE_CLOUD_SQL") == "true"
	cfg := Config{
		DBPort:      os.Getenv("DB_PORT"),
		DBUser:      os.Getenv("DB_USER"),
		DBPass:      os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		UseCloudSQL: useCloudSQL,
	}
	if useCloudSQL {
		cfg.DBHost = fmt.Sprintf("%s:%s:%s", os.Getenv("GCP_PROJECT_ID"), os.Getenv("DB_REGION"), os.Getenv("DB_INSTANCE"))
	} else {
		cfg.DBHost = os.Getenv("DB_HOST")
	}
	return cfg
}

func NewTestConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env ファイルが見つかりませんでした")
	}
	return Config{
		DBHost: os.Getenv("TEST_DB_HOST"),
		DBPort: os.Getenv("TEST_DB_PORT"),
		DBUser: os.Getenv("TEST_DB_USER"),
		DBPass: os.Getenv("TEST_DB_PASSWORD"),
		DBName: os.Getenv("TEST_DB_NAME"),
	}
}
