package dto

import "github.com/wasuwa/night-view-api/domain/model"

// NearestStation 最寄り駅
type NearestStation struct {
	RouteName              string
	StationName            string
	WalkingTimeFromStation int
}

// NearestStations 最寄り駅一覧
type NearestStations []NearestStation

// NewNearestStations 最寄り駅一覧を初期化する
func NewNearestStations(nearestStations []*model.NearestStation) NearestStations {
	results := make(NearestStations, len(nearestStations))
	for i, station := range nearestStations {
		results[i] = NearestStation{
			RouteName:              station.RouteName,
			StationName:            station.StationName,
			WalkingTimeFromStation: station.WalkingTimeFromStation,
		}
	}
	return results
}
