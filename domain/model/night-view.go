package model

// NightView 夜景
type NightView struct {
	ID         string
	Title      string
	PostCode   string
	Prefecture string
	City       string
	Address    string
	ImageURL   string
	Latitude   float64
	Longitude  float64
}
