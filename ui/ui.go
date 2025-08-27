package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/MUstinov77/weatherGo/env"
	"github.com/MUstinov77/weatherGo/utils"
)

func SetupUi(weatherResponse *utils.WeatherResponse, c *env.Config) {
	app := app.New()

	w := app.NewWindow("Forecast")

	var (
		timezoneLabel  = widget.NewLabel("")
		tempLabel      = widget.NewLabel("")
		feelsLikeLabel = widget.NewLabel("")
		humidityLabel  = widget.NewLabel("")
	)

	timezoneLabel.Text = fmt.Sprintf("Timezone %v", weatherResponse.Name)
	tempLabel.Text = fmt.Sprintf("Temperature: %.1f", weatherResponse.Main.Temp)
	feelsLikeLabel.Text = fmt.Sprintf("Feels like: %v", weatherResponse.Main.FeelsLike)
	humidityLabel.Text = fmt.Sprintf("Humidity: %v", weatherResponse.Main.Humidity)

	var vBox = container.NewVBox(
		timezoneLabel,
		tempLabel,
		feelsLikeLabel,
		humidityLabel,
	)
	vBox.Add(widget.NewButton("Close", func() {
		app.Quit()
	}))
	vBox.Add(widget.NewButton("Settings", func() {
		manageSettings(app, c)
	}))
	vBox.Add(widget.NewButton("change", func() {
		newButton := widget.NewLabel(fmt.Sprintf("new data - %v", "mdmcdc"))
		vBox.Add(newButton)
	}))
	w.SetContent(vBox)

	w.ShowAndRun()
}

func manageSettings(app fyne.App, c *env.Config) {
	w := app.NewWindow("Settings")

	var vBox = container.NewVBox(
		// TODO: add api urls from config
		widget.NewSelect([]string{c.Url, "nvvdkv"}, func(url string) {
			c.Url = url
		}),
		widget.NewButton("Close", func() {
			w.Close()
		}),
	)
	w.SetContent(vBox)
	w.Show()
}
