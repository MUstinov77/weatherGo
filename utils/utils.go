package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	urls "net/url"
)

func MakeResponse(url string, key string) (*WeatherResponse, error) {
	params := urls.Values{}
	params.Add("q", "Moscow")
	params.Add("appid", key)

	fullUrl := url + "?" + params.Encode()

	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Fatalf("Err till making request to api: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Api error: %s", body)
	}
	var weatherResponse WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		return nil, err
	}
	weatherResponse.Main.KelvinToCelsius()
	return &weatherResponse, nil
}

type WeatherResponse struct {
	Name string      `json:now`
	Main WeatherFact `json:main`
}

type WeatherFact struct {
	Temp       float64 `json:temp`
	Feels_Like float64 `json:feels_like`
	Humidity   int     `json:condition`
}

func (w *WeatherFact) FahrenheitToCelsius() {
	w.Temp = (w.Temp - 32) * 0.55
}

func (w *WeatherFact) KelvinToCelsius() {
	w.Temp = w.Temp - 273.15
	w.Feels_Like = w.Feels_Like - 273.15

}
