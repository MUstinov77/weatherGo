package main

import (
	"fmt"

	"github.com/MUstinov77/messanger/ui"
	"github.com/MUstinov77/messanger/utils"
)

func main() {
	// URL, к которому делаем запрос
	url := "https://api.openweathermap.org/data/2.5/weather"
	const apiKey = "589f351e2d6f2a3584e83f7879b34977"

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
