package datastore

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/wasuwa/night-view-api/config"
	"github.com/wasuwa/night-view-api/domain/model"
	"github.com/wasuwa/night-view-api/domain/repository"
)

type nightViewStore struct {
	db *bun.DB
}

// NewNightViewStore NightViewRepositoryの実装を初期化する
func NewNightViewStore(db *bun.DB) repository.NightViewRepository {
	return &nightViewStore{db: db}
}

// FindByID IDに紐づく夜景を取得する
func (ns *nightViewStore) FindByID(ctx context.Context, id string) (*model.NightView, error) {
	n := &model.NightView{ID: id}
	if err := ns.db.NewSelect().Model(n).WherePK().Scan(ctx); err != nil {
		return nil, err
	}
	return n, nil
}

// NewDB データベース接続を初期化する
func NewDB() *bun.DB {
	pgConn := pgdriver.NewConnector(
		pgdriver.WithAddr("localhost:5437"),
		pgdriver.WithUser(config.PostgresUser),
		pgdriver.WithPassword(config.PostgresPassword),
		pgdriver.WithDatabase(config.PostgresDB),
		pgdriver.WithApplicationName("myapp"),
	)
	db := bun.NewDB(sql.OpenDB(pgConn), pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))
	return db
}
