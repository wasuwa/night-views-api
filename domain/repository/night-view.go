package repository

import (
	"context"

	"github.com/wasuwa/night-view-api/domain/model"
)

// NightViewRepository 夜景のリポジトリ
type NightViewRepository interface {
	FindByID(ctx context.Context, id string) (*model.NightView, error)
}
