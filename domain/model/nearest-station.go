package model

// NearestStation 最寄り駅
type NearestStation struct {
	ID                     string
	NightViewID            string
	RouteName              string
	StationName            string
	WalkingTimeFromStation int
}
