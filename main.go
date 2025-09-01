package main

import (
	"fmt"

	"github.com/MUstinov77/weatherGo/app"
	"github.com/MUstinov77/weatherGo/config"
)

/*
	func main() {
		// Parsing env file and build map[nameOfKey]=[keyVariable]
		envMap, err := config.ParseEnvFile(".env")
		if err != nil {
			fmt.Println("Tvoi algos ne pashet")
		}
		// Build config and load-in the env-variables
		config := config.NewConfig()
		config.LoadConfig(envMap)
		// TODO: run prog in cycle to fetch the data from api and update ui
		// Build request to open.weather.api and parse the response
		weatherResponse, err := utils.MakeResponse(config.Url, config.ApiKey)
		if err != nil {
			fmt.Println("Error during fetch api")
			fmt.Println(err)
			return
		}

		// Set up ui
		ui.SetupUi(weatherResponse, config)
	}
*/
func main() {
	envMap, err := config.ParseEnvFile(".env")
	if err != nil {
		fmt.Println(err)
	}

	weatherApp := app.NewWeatherApp()
	weatherApp.Config.LoadConfig(envMap)

	weatherApp.Run()

}
