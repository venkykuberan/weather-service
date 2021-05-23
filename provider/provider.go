package provider

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Weatherdto struct {
	Lat            float64
	Lon            float64
	Timezone       string
	TimezoneOffset int
	Temp           float64
	FeelsLike      float64
	Description    string
	Alerts         []Alert
}

func GetWeatherInfo(url string, api_key string, lat string, lon string) (*Weatherdto, error) {
	var client http.Client
	u, err := addQueryParams(url, map[string]string{
		"lat":     lat,
		"lon":     lon,
		"exclude": "hourly,daily,minutely",
		"units":   "imperial",
		"appid":   api_key,
	})
	if err != nil {
		return nil, errors.New("failing to create request queury params: " + err.Error())
	}

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, errors.New("failing to create http request : " + err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New("request to api failed : " + err.Error())
	}
	defer res.Body.Close()
	weatherInfo, err := getParseResponse(res)
	if err != nil {
		return nil, errors.New("parsing api response failed : " + err.Error())
	}
	return weatherInfo, nil
}

func addQueryParams(rawurl string, params map[string]string) (string, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func getParseResponse(res *http.Response) (*Weatherdto, error) {
	if res.StatusCode > 200 && res.StatusCode <= 299 {
		return nil, errors.New("api responded status code  : " + string(res.StatusCode))
	}

	slurp, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(slurp))

	if err != nil {
		return nil, err
	}

	var weatherSrc WeatherSource
	weatherdto := Weatherdto{}

	err = json.Unmarshal(slurp, &weatherSrc)
	if err != nil {
		return nil, err
	}

	//mapping source data to dto
	weatherdto.Lat = weatherSrc.Lat
	weatherdto.Lon = weatherSrc.Lon
	weatherdto.FeelsLike = weatherSrc.Current.FeelsLike
	weatherdto.Description = weatherSrc.Current.Weather[0].Description // ToDo need enhancement
	weatherdto.Temp = weatherSrc.Current.Temp
	weatherdto.Timezone = weatherSrc.Timezone
	weatherdto.Alerts = weatherSrc.Alerts

	return &weatherdto, nil
}
