package main

import (
	"fmt"

	"github.com/MUstinov77/weatherGo/env"
	"github.com/MUstinov77/weatherGo/ui"
	"github.com/MUstinov77/weatherGo/utils"
)

func main() {
	// URL, к которому делаем запрос
	url := "https://api.openweathermap.org/data/2.5/weather"
	envMap, err := env.ParseEnvFile(".env")
	if err != nil {
		fmt.Println("Tvoi algos ne pashet")
	}
	apiKey := envMap["API_KEY"]
	fmt.Println(envMap, apiKey)

	// Создаем новый GET-запрос
	weatherResponse, err := utils.MakeResponse(url, apiKey)
	if err != nil {
		fmt.Println("Error during fetch api")
		fmt.Println(err)
		return
	}

	// Читаем тело ответа
	ui.SetupUi(weatherResponse)
}
