package datastore

import (
	"context"

	"github.com/uptrace/bun"
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
	// pgConn := pgdriver.NewConnector(
	// 	pgdriver.WithNetwork("tcp"),
	// 	pgdriver.WithAddr("localhost:5437"),
	// 	pgdriver.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
	// 	pgdriver.WithUser("test"),
	// 	pgdriver.WithPassword("test"),
	// 	pgdriver.WithDatabase("test"),
	// 	pgdriver.WithApplicationName("myapp"),
	// 	pgdriver.WithTimeout(5 * time.Second),
	// 	pgdriver.WithDialTimeout(5 * time.Second),
	// 	pgdriver.WithReadTimeout(5 * time.Second),
	// 	pgdriver.WithWriteTimeout(5 * time.Second),
	// 	pgdriver.WithConnParams(map[string]interface{}{
	// 		"search_path": "my_search_path",
	// 	}),
	// )
	// return bun.NewDB(sql.OpenDB(pgConn), pgdialect.New())
	return &bun.DB{}
}
