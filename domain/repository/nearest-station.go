package repository

import (
	"context"

	"github.com/wasuwa/night-view-api/domain/model"
)

// NearestStationRepository 最寄り駅のリポジトリ
type NearestStationRepository interface {
	FindByNightViewID(ctx context.Context, nightViewID string) ([]*model.NearestStation, error)
}
