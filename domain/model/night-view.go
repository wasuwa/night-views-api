package model

// NightView 夜景
type NightView struct {
	ID              string
	Title           string
	ImageURL        string
	PostCode        string
	Address         string
	Latitude        float64
	Longitude       float64
	NearestStations []*NearestStation
}
