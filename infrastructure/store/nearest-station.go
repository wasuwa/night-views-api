package store

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/wasuwa/night-view-api/domain/model"
	"github.com/wasuwa/night-view-api/domain/repository"
)

// NearestStation 最寄り駅
type NearestStation struct {
	bun.BaseModel          `bun:"table:nearest_stations"`
	ID                     string `bun:"id,pk"`
	NightViewID            string `bun:"night_view_id"`
	RouteName              string
	StationName            string
	WalkingTimeFromStation int
}

// domainModel ドメインモデルに変換する
func (n NearestStation) domainModel() *model.NearestStation {
	return &model.NearestStation{
		ID:                     n.ID,
		NightViewID:            n.NightViewID,
		RouteName:              n.RouteName,
		StationName:            n.StationName,
		WalkingTimeFromStation: n.WalkingTimeFromStation,
	}
}

// NearestStations 最寄り駅一覧
type NearestStations []NearestStation

// domainModel ドメインモデルに変換する
func (n NearestStations) domainModel() []*model.NearestStation {
	result := make([]*model.NearestStation, len(n))
	for i, station := range n {
		result[i] = station.domainModel()
	}
	return result
}

type nearestStationStore struct {
	db *bun.DB
}

// NewNearestStationStore NearestStationRepositoryの実装を初期化する
func NewNearestStationStore(db *bun.DB) repository.NearestStationRepository {
	return &nearestStationStore{db: db}
}

// FindByNightViewID 夜景IDに紐づく最寄り駅を取得する
func (ns *nearestStationStore) FindByNightViewID(ctx context.Context, nightViewID string) ([]*model.NearestStation, error) {
	var stations NearestStations
	err := ns.db.NewSelect().
		Column("id", "night_view_id", "route_name", "station_name", "walking_time_from_station").
		Model(&stations).
		Where("night_view_id = ?", nightViewID).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return stations.domainModel(), nil
}
