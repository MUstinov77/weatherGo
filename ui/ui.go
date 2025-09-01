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

func SetupUi(weatherResponse *utils.WeatherResponse, c *config.Config) {
	app := app.New()

	w := app.NewWindow("Forecast")

	var (
		timezoneLabel  = widget.NewLabel("")
		tempLabel      = widget.NewLabel("")
		feelsLikeLabel = widget.NewLabel("")
		humidityLabel  = widget.NewLabel("")
	)
	// var noneLabel *widget.Label = widget.NewLabel("")

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

func manageSettings(app fyne.App, c *config.Config) {
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
	// TODO: add settings func in new ui struct
	/* vBox.Add(widget.NewButton("Settings", func() {
		manageSettings(app, c)
	}))*/

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
	ui.feelsLikeLabel.SetText(fmt.Sprintf("Feels like: %v", resp.Main.FeelsLike))
	ui.humidityLabel.SetText(fmt.Sprintf("Humidity: %v", resp.Main.Humidity))
	ui.lastUpdated.SetText(fmt.Sprintf("Last updated: %v", time.Now().Format()))
}
