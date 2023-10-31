package models

import (
	"fmt"
	"math/rand"
	"time"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Vehicle struct {
    parking *Parking
	I               int
	ParkingSpace int
	skin			*canvas.Image
}

func NewVehicle(p *Parking, s *canvas.Image) *Vehicle {
	return &Vehicle{
        parking: p,
		skin: s,
	}
}


func (v *Vehicle) RunVehicle() {
	v.parking.Space <- true
	v.parking.mutex.Lock()	
		for i := 0; i < len(v.parking.ParkingSpaces); i++ {
			if !v.parking.ParkingSpaces[i].occupied{
				v.skin.Move(fyne.NewPos(v.parking.ParkingSpaces[i].x, v.parking.ParkingSpaces[i].y))
				v.skin.Refresh()
				v.ParkingSpace = i
				v.parking.ParkingSpaces[i].occupied = true
				break
			}
		}
	
		fmt.Println("vehicle ", v.I, " enters")
		time.Sleep(1 *time.Millisecond)
		v.parking.mutex.Unlock()

		randomSleepSeconds := rand.Intn(20) + 6
		time.Sleep(time.Duration(randomSleepSeconds) * time.Second)

	v.parking.mutex.Lock()
		<- v.parking.Space
		v.parking.ParkingSpaces[v.ParkingSpace].occupied = false
		v.skin.Move(fyne.NewPos( 460,45 ))
		fmt.Println("vehicle ", v.I, "exit")
		time.Sleep(1 *time.Millisecond)
        v.parking.mutex.Unlock()
}



func GenerateVehicle(n int, parking *Parking) {
	parking.Space <- true
	for i := 0; i < n; i++ {
		randomVehicleNumber := rand.Intn(7) + 1
		vehicleImageName := fmt.Sprintf("./assets/vehicle%d.png", randomVehicleNumber)

        VehicleImg := canvas.NewImageFromURI(storage.NewFileURI(vehicleImageName))
		VehicleImg.Resize(fyne.NewSize(70,120))
		VehicleImg.Move(fyne.NewPos(460, 650))

		NewVehicle := NewVehicle(parking, VehicleImg)
		NewVehicle.I = i + 1

		parking.DrawVehicle <- VehicleImg
		time.Sleep(time.Millisecond*1)
		go NewVehicle.RunVehicle()
		time.Sleep(time.Duration(rand.ExpFloat64() * float64(time.Second)))
	}
}
