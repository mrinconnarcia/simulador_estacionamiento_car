package controllers

import (
	"estacionamiento/models"
	"estacionamiento/views"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
)

func tiempoLlegadaPoisson(lambda float64) time.Duration {
	// generar un número aleatorio uniforme entre 0 y 1
	U := rand.Float64()

	// calcular el tiempo de llegada basado en la distribución Poisson
	tiempo := -math.Log(U) / lambda

	// convertir el tiempo a duración para usarlo en time.Sleep
	tiempoDuracion := time.Duration(tiempo * float64(time.Second))

	return tiempoDuracion
}

func GenerarCarros(cantidad int, estacionamiento *models.Parking, ventana fyne.Window) {
	var wg sync.WaitGroup
	estacionamiento.Cajones <- true

	for i := 1; i <= cantidad; i++ {
		nuevoCarro := CrearCarro(i, estacionamiento)
		estacionamiento.Posiciones <- nuevoCarro.Rectangulo
		wg.Add(1)
		go func() {
			MoverCarroConAnimacion(nuevoCarro, 750, 600) // animación hacia la posición de espera
			ConducirCarro(nuevoCarro)
			wg.Done()
		}()

		tiempoEspera := tiempoLlegadaPoisson(1.0) // lambda valor
		time.Sleep(tiempoEspera)
	}
	wg.Wait()
	dialog.ShowInformation("Información", "Todos los carros han sido generados.", ventana)
}


func CrearCarro(id int, p *models.Parking) *models.Carro {
	rectanguloCarro := canvas.NewRectangle(views.GenerarColorAleatorio())
	rectanguloCarro.Resize(fyne.NewSize(40, 80))
	rectanguloCarro.Move(fyne.NewPos(750, 1000)) // posición inicial fuera de la pantalla

	nuevoCarro := models.NuevoCarro(id, p, rectanguloCarro)
	nuevoCarro.ID = id

	fmt.Printf("El carro %d está llegando al estacionamiento.\n", nuevoCarro.ID)

	return nuevoCarro
}

func MoverCarroConAnimacion(c *models.Carro, esperaX, esperaY float32) {
	posXInicial := c.Rectangulo.Position().X
	posYInicial := c.Rectangulo.Position().Y

	// Número de pasos para la animación
	pasos := 100

	// Cálculo de desplazamiento por paso
	dx := (esperaX - posXInicial) / float32(pasos)
	dy := (esperaY - posYInicial) / float32(pasos)

	// Realizar la animación
	for i := 1; i <= pasos; i++ {
		posX := posXInicial + dx*float32(i)
		posY := posYInicial + dy*float32(i)
		c.Rectangulo.Move(fyne.NewPos(posX, posY))
		time.Sleep(time.Millisecond * 10)
	}
}


// func MoverCarro(c *models.Carro, nuevoX, nuevoY float32) {
// 	c.MoverA(nuevoX, nuevoY)
// }

func ConducirCarro(c *models.Carro) {
	c.Conducir()
}
