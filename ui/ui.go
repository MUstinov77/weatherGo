package ui

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/MUstinov77/weatherGo/config"
	"github.com/MUstinov77/weatherGo/utils"
)

type UI struct {
	app            fyne.App
	Window         fyne.Window
	vBox           *fyne.Container
	timezoneLabel  *widget.Label
	tempLabel      *widget.Label
	feelsLikeLabel *widget.Label
	humidityLabel  *widget.Label
	lastUpdated    *widget.Label
}

func NewUi() *UI {
	app := app.New()

	window := app.NewWindow("Forecast")

	timezoneLabel := widget.NewLabel("Загрузка...")
	tempLabel := widget.NewLabel("Загрузка...")
	feelsLikeLabel := widget.NewLabel("Загрузка...")
	humidityLabel := widget.NewLabel("Загрузка...")
	lastUpdated := widget.NewLabel("Загрузка...")

	vBox := container.NewVBox(
		timezoneLabel,
		tempLabel,
		feelsLikeLabel,
		humidityLabel,
		lastUpdated,
	)

	vBox.Add(widget.NewButton("Close", func() {
		app.Quit()
	}))

	window.SetContent(vBox)

	return &UI{
		app:            app,
		Window:         window,
		vBox:           vBox,
		timezoneLabel:  timezoneLabel,
		tempLabel:      tempLabel,
		feelsLikeLabel: feelsLikeLabel,
		humidityLabel:  humidityLabel,
		lastUpdated:    lastUpdated,
	}
}

func (ui *UI) UpdateUI(config *config.Config) {
	resp, err := utils.MakeResponse(config.Url, config.ApiKey)
	if err != nil {
		fmt.Println(err)
	}

	ui.timezoneLabel.SetText(fmt.Sprintf("Timezone: %v", resp.Name))
	ui.tempLabel.SetText(fmt.Sprintf("Temperature: %.1f", resp.Main.Temp))
	ui.feelsLikeLabel.SetText(fmt.Sprintf("Feels like: %.1f", resp.Main.Feels_Like))
	ui.humidityLabel.SetText(fmt.Sprintf("Humidity: %v", resp.Main.Humidity))
	ui.lastUpdated.SetText(fmt.Sprintf("Last updated: %v", time.Now().Format(time.TimeOnly)))
}
