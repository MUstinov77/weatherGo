package main

import (
	"fmt"

	"github.com/MUstinov77/weatherGo/env"
	"github.com/MUstinov77/weatherGo/ui"
	"github.com/MUstinov77/weatherGo/utils"
)

func main() {
	// URL, к которому делаем запрос
	envMap, err := env.ParseEnvFile(".env")
	if err != nil {
		fmt.Println("Tvoi algos ne pashet")
	}
	config := env.Config{}
	config.LoadConfig(envMap)

	// Создаем новый GET-запрос
	weatherResponse, err := utils.MakeResponse(config.Url, config.ApiKey)
	if err != nil {
		fmt.Println("Error during fetch api")
		fmt.Println(err)
		return
	}

	// Читаем тело ответа
	ui.SetupUi(weatherResponse)
}
