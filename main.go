package main

import (
	"fmt"

	"github.com/MUstinov77/weatherGo/app"
	"github.com/MUstinov77/weatherGo/config"
)

func main() {
	envMap, err := config.ParseEnvFile(".env")
	if err != nil {
		fmt.Println(err)
	}

	weatherApp := app.NewWeatherApp()
	weatherApp.Config.LoadConfig(envMap)

	weatherApp.Run()

}
