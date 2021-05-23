package main

import (
	"log"
	"net/http"
	svc "weatherinfo/weatherforecast/service"
)

func main() {
	http.HandleFunc("/weather/", svc.WeatherHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
