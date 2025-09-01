package app

import (
	"time"

	"fyne.io/fyne/v2"
	"github.com/MUstinov77/weatherGo/ui"

	"github.com/MUstinov77/weatherGo/config"
)

type WeatherApp struct {
	Config *config.Config
	UI     *ui.UI
}

func (wa *WeatherApp) Run() {

	wa.UI.UpdateUI(wa.Config)

	go wa.startAutoUpdate()

	wa.UI.Window.ShowAndRun()
}

func (wa *WeatherApp) startAutoUpdate() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	wa.UI.UpdateUI(wa.Config)

	for {
		select {
		case <-ticker.C:
			fyne.Do(func() {
				wa.UI.UpdateUI(wa.Config)
			})
		}
	}

}

func NewWeatherApp() *WeatherApp {
	return &WeatherApp{
		Config: config.NewConfig(),
		UI:     ui.NewUi(),
	}
}
