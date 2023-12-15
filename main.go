package main

import (
	"estacionamiento/scenes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	window := app.NewWindow("Simulador de estacionamiento de carros - Golang")

	window.Resize(fyne.NewSize(990, 740))
	scenes.NewParkingScene(window)

	window.ShowAndRun()
}
