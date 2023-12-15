// views/parking_view.go
package views

import (
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

type ParkingView struct {
	Container *fyne.Container
}

func NewParkingView() *ParkingView {
	view := &ParkingView{}
	view.Generate()
	return view
}

func (pv *ParkingView) Generate() {
	backgroundImage := pv.createBackgroundImage()
	pv.Container = container.NewWithoutLayout(backgroundImage)
}

func (pv *ParkingView) createBackgroundImage() *canvas.Image {
	backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/estacionamiento_bg.png"))
	backgroundImage.Resize(fyne.NewSize(990, 740))
	backgroundImage.Move(fyne.NewPos(0, 0))
	return backgroundImage
}


func GenerarColorAleatorio() color.RGBA {
	return color.RGBA{R: uint8(rand.Intn(256)), G: uint8(rand.Intn(256)), B: uint8(rand.Intn(256)), A: 255}
}
