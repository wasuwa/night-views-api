package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	// PostgresUser PostgresSQLのユーザー名
	PostgresUser string
	// PostgresPassword PostgresSQLのパスワード
	PostgresPassword string
	// PostgresDB PostgresSQLのデータベース名
	PostgresDB string
)

// LoadEnv 環境変数を読み込む
func LoadEnv() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	PostgresUser = os.Getenv("POSTGRES_USER")
	PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	PostgresDB = os.Getenv("POSTGRES_DB")
	return nil
}
