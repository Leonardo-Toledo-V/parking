package models

type Vehicle struct {
    number int
}

func NewVehicle(n int) *Vehicle {
    return &Vehicle{
        number: n,
    }
}

func (v *Vehicle) Number() int {
    return v.number
}
