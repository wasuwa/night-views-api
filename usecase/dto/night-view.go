package dto

// NightView 夜景
type NightView struct {
	ID              string
	Title           string
	ImageURL        string
	PostCode        string
	Address         string
	Latitude        float64
	Longitude       float64
	NearestStations NearestStations
}

// NewNightView 夜景を初期化する
func NewNightView(id, title, imageURL, postCode, address string, latitude, longitude float64, nearestStations NearestStations) *NightView {
	return &NightView{
		ID:              id,
		Title:           title,
		ImageURL:        imageURL,
		PostCode:        postCode,
		Address:         address,
		Latitude:        latitude,
		Longitude:       longitude,
		NearestStations: nearestStations,
	}
}
