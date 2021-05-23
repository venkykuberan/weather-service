package service

import "weatherinfo/weatherforecast/provider"

type CurrentWeather struct {
	Lat         float64
	Lon         float64
	Temperature float64
	FeelsLike   float64
	Description string
	TimeZone    string
	Alerts      []provider.Alert
}
