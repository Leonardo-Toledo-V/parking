package controllers

import (
    "fmt"
    "parking/models"
    "math/rand"
	"math"
    "time"
)

type VehicleController struct{}

func NewVehicleController() *VehicleController {
    return &VehicleController{}
}

func (vc *VehicleController) CreateVehicles(numVehicles int, poissonLambda float64) {
    rand.Seed(time.Now().UnixNano())

    for i := 1; i <= numVehicles; i++ {
        waitTime := -math.Log(1-rand.Float64()) / poissonLambda
        time.Sleep(time.Duration(waitTime * float64(time.Second)))

        vehicle := models.NewVehicle(i)
        fmt.Printf("Vehículo %d creado.\n", vehicle.Number())
    }
}
