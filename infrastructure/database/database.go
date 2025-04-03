package database

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/wasuwa/night-view-api/config"
)

// NewDB データベース接続を初期化する
func NewDB() *bun.DB {
	pgConn := pgdriver.NewConnector(
		pgdriver.WithAddr(config.PostgresHost),
		pgdriver.WithUser(config.PostgresUser),
		pgdriver.WithPassword(config.PostgresPassword),
		pgdriver.WithDatabase(config.PostgresDB),
		pgdriver.WithInsecure(true),
	)
	return bun.NewDB(sql.OpenDB(pgConn), pgdialect.New())
}
