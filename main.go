package main

import (
    "parking/controllers"
)

func main() {
    vehicleController := controllers.NewVehicleController()
    numVehicles := 100
    poissonLambda := 1.0

    vehicleController.CreateVehicles(numVehicles, poissonLambda)
}
