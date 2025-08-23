package main

import (
	"fmt"

	"github.com/MUstinov77/weatherGo/env"
	"github.com/MUstinov77/weatherGo/ui"
	"github.com/MUstinov77/weatherGo/utils"
)

func main() {
	// Parsing env file and build map[nameOfKey]=[keyVariable]
	envMap, err := env.ParseEnvFile(".env")
	if err != nil {
		fmt.Println("Tvoi algos ne pashet")
	}
	// Build config and load-in the env-variables
	config := env.Config{}
	config.LoadConfig(envMap)

	// Build request to the default weather api and parse the response
	weatherResponse, err := utils.MakeResponse(config.Url, config.ApiKey)
	if err != nil {
		fmt.Println("Error during fetch api")
		fmt.Println(err)
		return
	}

	// Set up ui
	ui.SetupUi(weatherResponse, &config)
}
