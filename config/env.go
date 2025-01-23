package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	// PostgresUser Postgresのユーザー名
	PostgresUser string
	// PostgresPassword Postgresのパスワード
	PostgresPassword string
)

// LoadEnv 環境変数を読み込む
func LoadEnv() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	PostgresUser = os.Getenv("POSTGRES_USER")
	PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	return nil
}
