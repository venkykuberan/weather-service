package service

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"weatherinfo/weatherforecast/provider"
)

var URL = "https://api.openweathermap.org/data/2.5/onecall"

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")

	if len(lat) == 0 || len(lon) == 0 {
		w.WriteHeader(401)
		w.Write([]byte("bad url"))
	} else {
		weatherDto, err := provider.GetWeatherInfo(URL, os.Getenv("API_KEY"), lat, lon)
		if err != nil {
			log.Fatal(errors.New("Error fetching current weather information"))
		}

		weatherInfo := CurrentWeather{}

		// map the dto to service response
		weatherInfo.Lat = weatherDto.Lat
		weatherInfo.Lon = weatherDto.Lon
		weatherInfo.TimeZone = weatherDto.Timezone
		weatherInfo.Description = weatherDto.Description
		weatherInfo.Temperature = weatherDto.Temp
		weatherInfo.FeelsLike = weatherDto.FeelsLike
		if len(weatherDto.Alerts) > 0 {
			weatherInfo.Alerts = weatherDto.Alerts
		}

		bytes, err := json.Marshal(weatherInfo)

		if err != nil {
			log.Fatal(err.Error())
		}

		w.Write(bytes)
	}

}
