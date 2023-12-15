// scenes/parking_scene.go
package scenes

import (
	"estacionamiento/controllers"
	"estacionamiento/models"
	"estacionamiento/views"

	"fyne.io/fyne/v2"
)

type ParkingScene struct {
	Window fyne.Window
	View   *views.ParkingView
}

func NewParkingScene(window fyne.Window) *ParkingScene {
	scene := &ParkingScene{Window: window}
	scene.View = views.NewParkingView()
	scene.Window.SetContent(scene.View.Container)
	scene.Init()
	return scene
}

func (ps *ParkingScene) Init() {
	parking := models.NuevoParking()
	go controllers.GenerarCarros(100, parking, ps.Window)
	go ps.showCarsPosition(parking)
}

func (ps *ParkingScene) showCarsPosition(parking *models.Parking) {
	for {
		obj := <-parking.Posiciones
		ps.View.Container.Add(obj)
		ps.Window.Canvas().Refresh(ps.View.Container)
	}
}
