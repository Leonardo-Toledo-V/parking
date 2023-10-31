package models

import (
	"sync"
	"fyne.io/fyne/v2/canvas"
)

type Slot struct{
	x float32
	y float32
	occupied bool
}

type Parking struct {
	Space chan bool
	DrawVehicle chan *canvas.Image
	mutex sync.Mutex
	ParkingSpaces []Slot
}


func NewParking(nS int) *Parking {
	return &Parking{
		Space: make(chan bool, nS+1),
		DrawVehicle: make(chan *canvas.Image,1),
		ParkingSpaces: []Slot{
			{x: 60, y: 70, occupied:false},
			{x: 130, y: 70, occupied:false},
			{x: 200, y: 70, occupied:false},
			{x: 275, y: 70, occupied:false},
			{x: 345, y: 70, occupied:false},
			{x: 415, y: 70, occupied:false},
			{x: 485, y: 70, occupied:false},
			{x: 560, y: 70, occupied:false},
			{x: 635, y: 70, occupied:false},
			{x: 700, y: 70, occupied:false},

			{x: 60, y: 230, occupied:false},
			{x: 130, y: 230, occupied:false},
			{x: 200, y: 230, occupied:false},
			{x: 275, y: 230, occupied:false},
			{x: 345, y: 230, occupied:false},
			{x: 415, y: 230, occupied:false},
			{x: 485, y: 230, occupied:false},
			{x: 560, y: 230, occupied:false},
			{x: 635, y: 230, occupied:false},
			{x: 700, y: 230, occupied:false},
		},
	}
}