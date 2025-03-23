package store

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/wasuwa/night-view-api/config"
	"github.com/wasuwa/night-view-api/domain/model"
	"github.com/wasuwa/night-view-api/domain/repository"
)

// NightView 夜景
type NightView struct {
	bun.BaseModel `bun:"table:night_views"`
	ID            string `bun:"pk"`
	Title         string
	PostCode      string
	Prefecture    string
	City          string
	Address       string
	ImageURL      string
	Location      string
}

type nightViewStore struct {
	db *bun.DB
}

// NewNightViewStore NightViewRepositoryの実装を初期化する
func NewNightViewStore(db *bun.DB) repository.NightViewRepository {
	return &nightViewStore{db: db}
}

// FindByID IDに紐づく夜景を取得する
func (ns *nightViewStore) FindByID(ctx context.Context, id string) (*model.NightView, error) {
	n := &NightView{ID: id}
	if err := ns.db.NewSelect().Model(n).WherePK().Scan(ctx); err != nil {
		return nil, err
	}
	return &model.NightView{
		ID:         n.ID,
		Title:      n.Title,
		PostCode:   n.PostCode,
		Prefecture: n.Prefecture,
		City:       n.City,
		Address:    n.Address,
		ImageURL:   n.ImageURL,
		// Latitude:    n.Latitude,
		// Longitude:   n.Longitude,
	}, nil
}

// NewDB データベース接続を初期化する
func NewDB() *bun.DB {
	pgConn := pgdriver.NewConnector(
		pgdriver.WithAddr(config.PostgresHost),
		pgdriver.WithUser(config.PostgresUser),
		pgdriver.WithPassword(config.PostgresPassword),
		pgdriver.WithDatabase(config.PostgresDB),
		pgdriver.WithInsecure(true),
	)
	db := bun.NewDB(sql.OpenDB(pgConn), pgdialect.New())
	return db
}
