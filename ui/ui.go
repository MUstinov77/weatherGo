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

	var vBox = container.NewVBox(
		widget.NewLabel(fmt.Sprintf("Timezone %v", weatherResponse.Name)),
		widget.NewLabel(fmt.Sprintf("Temperature: %.1f", weatherResponse.Main.Temp)),
		widget.NewLabel(fmt.Sprintf("Feels like: %v", weatherResponse.Main.FeelsLike)),
		widget.NewLabel(fmt.Sprintf("Humidity: %v", weatherResponse.Main.Humidity)),
	)
	vBox.Add(widget.NewButton("Close", func() {
		app.Quit()
	}))
	vBox.Add(widget.NewButton("Settings", func() {
		manageSettings(app, c)
	}))

	w.SetContent(vBox)

	w.ShowAndRun()
}

func manageSettings(app fyne.App, c *env.Config) {
	w := app.NewWindow("Settings")

	var vBox = container.NewVBox(
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
