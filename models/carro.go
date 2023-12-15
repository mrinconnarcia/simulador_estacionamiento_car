// models/carro.go
package models

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Carro struct {
	ID         int
	Parking    *Parking
	Posicion   int
	Rectangulo *canvas.Rectangle
}

func NuevoCarro(id int, p *Parking, r *canvas.Rectangle) *Carro {
	return &Carro{
		ID:         id,
		Parking:    p,
		Rectangulo: r,
	}
}

func (c *Carro) MoverA(nuevoX, nuevoY float32) {
	const pasos = 70
	dx := (nuevoX - c.Rectangulo.Position().X) / pasos
	dy := (nuevoY - c.Rectangulo.Position().Y) / pasos
	for i := 1; i <= pasos; i++ {
		c.Rectangulo.Move(fyne.NewPos(c.Rectangulo.Position().X+dx, c.Rectangulo.Position().Y+dy))
		c.Rectangulo.Refresh()
		time.Sleep(time.Millisecond * 7)
	}
}

func (c *Carro) IngresarEstacionamiento() {
	c.Parking.Cajones <- true
	// fmt.Print("espacios ocupados: ", c.Parking.EspaciosOcupados , "\n")
	fmt.Printf("El carro %d está intentando ingresar al estacionamiento.\n", c.ID)
	c.Parking.PuertaEntrada <- true
	c.Parking.M.Lock()
	bestSpace := -1
	bestDistance := float32(math.MaxFloat32)

	for bestSpace == -1 {
		for i := 0; i < len(c.Parking.Espacios); i++ {
			if c.Parking.Espacios[i].Disponible {
				distance := distanciaEntreDosPuntos(c.Rectangulo.Position().X, c.Rectangulo.Position().Y, c.Parking.Espacios[i].X, c.Parking.Espacios[i].Y)
				if distance < bestDistance {
					bestDistance = distance
					bestSpace = i
				}
			}
		}

		if bestSpace == -1 {
			// si no hay espacio disponible, esperar un breve momento y volver a verificar
			c.Parking.M.Unlock()
			time.Sleep(500 * time.Millisecond)
			c.Parking.M.Lock()
		}
	}

	// c.MoverA(c.Parking.Espacios[bestSpace].X, c.Parking.Espacios[bestSpace].Y)
	c.MoverA(c.Rectangulo.Position().X, c.Rectangulo.Position().Y-160)
	c.MoverA(c.Rectangulo.Position().X+90, c.Rectangulo.Position().Y)
	c.MoverA(c.Rectangulo.Position().X, c.Rectangulo.Position().Y-250)

	// movimiento final hacia la posición del espacio disponible
	espacioX := c.Parking.Espacios[bestSpace].X
	espacioY := c.Parking.Espacios[bestSpace].Y
	c.MoverA(espacioX, espacioY)
	c.Posicion = bestSpace
	c.Parking.Espacios[bestSpace].Disponible = false
	c.Parking.EspaciosOcupados++
	fmt.Printf("El carro %d ha ingresado al estacionamiento y ha ocupado el espacio %d.\n", c.ID, bestSpace)
	time.Sleep(10 * time.Millisecond)
	c.Parking.M.Unlock()
	<-c.Parking.PuertaEntrada
}

// calcular la distancia entre dos puntos
func distanciaEntreDosPuntos(x1, y1, x2, y2 float32) float32 {
	return float32(math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))))
}

func (c *Carro) Esperar() {
	espera := rand.Intn(30) + 10
	// espera := rand.Intn(5) + 1
	time.Sleep(time.Duration(espera) * time.Second)
}

func (c *Carro) SalirEstacionamiento() {
	c.Parking.PuertaSalida <- true
	c.Parking.M.Lock()
	<-c.Parking.Cajones // Semaforo disponible
	c.Parking.Espacios[c.Posicion].Disponible = true
	c.Parking.EspaciosOcupados--
	fmt.Printf("Carro %d está saliendo del estacionamiento, del espacio %d.\n", c.ID, c.Posicion)

	//movimientos para salir del espacio hacia la entrada del estacionamiento
	c.MoverA(c.Rectangulo.Position().X, c.Rectangulo.Position().Y-100)
	c.MoverA(900, 250)
	// movimiento final hacia la posición de salida
	c.MoverA(900, 600)

	// tiempo de permanencia entre 0 y 2 segundos
	tiempoPermanencia := time.Duration(rand.Intn(2)) * time.Second
	time.Sleep(tiempoPermanencia)
	c.Parking.M.Unlock()
	<-c.Parking.PuertaSalida // espera confirmacion a la puerta de salida
	c.MoverA(900, 1000)
}

func (c *Carro) Conducir() {
	c.IngresarEstacionamiento()
	c.Esperar()
	c.SalirEstacionamiento()
}
