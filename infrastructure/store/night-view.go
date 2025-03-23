package store

import (
	"context"
	"database/sql"
	"errors"

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
	ID            string `bun:"id,pk"`
	Title         string
	ImageURL      string
	PostCode      string
	Address       string
	Longitude     float64
	Latitude      float64
}

// domainModel ドメインモデルに変換する
func (n NightView) domainModel() *model.NightView {
	return &model.NightView{
		ID:        n.ID,
		Title:     n.Title,
		ImageURL:  n.ImageURL,
		PostCode:  n.PostCode,
		Address:   n.Address,
		Longitude: n.Longitude,
		Latitude:  n.Latitude,
	}
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
	err := ns.db.NewSelect().
		Column("id", "title", "image_url", "post_code", "address").
		ColumnExpr("extensions.st_x(location::extensions.geometry) AS longitude").
		ColumnExpr("extensions.st_y(location::extensions.geometry) AS latitude").
		Model(n).
		WherePK().
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return n.domainModel(), nil
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
	return bun.NewDB(sql.OpenDB(pgConn), pgdialect.New())
}
