// models/parking.go
package models

import (
	"sync"

	"fyne.io/fyne/v2"
)

type Espacio struct {
	X          float32
	Y          float32
	Disponible bool
}

type CanvasObject = fyne.CanvasObject

type Parking struct {
	Cajones          chan bool
	Posiciones       chan CanvasObject
	M                sync.Mutex
	Espacios         []Espacio
	PuertaEntrada    chan bool
	PuertaSalida     chan bool
	EspaciosOcupados int
}

func NuevoParking() *Parking {
	const numColumnas = 20

	return &Parking{
		Cajones:          make(chan bool, numColumnas+1),
		Posiciones:       make(chan CanvasObject, 100),
		PuertaEntrada:    make(chan bool, 1),
		PuertaSalida:     make(chan bool, 1),
		Espacios:         inicializarEspacios(),
		EspaciosOcupados: 0,
	}
}

func inicializarEspacios() []Espacio {
	const (
		numFilas    = 2
		numColumnas = 10
	)

	var espacios []Espacio
	for fila := 0; fila < numFilas; fila++ {
		for columna := 0; columna < numColumnas; columna++ {
			x := float32(70 + columna*70)
			y := float32(100 + fila*220)
			espacio := Espacio{X: x, Y: y, Disponible: true}
			espacios = append(espacios, espacio)
		}
	}
	return espacios
}
