package main

import "fmt"

type Vihicle interface {
	Start()
	Stop()
}

type Car struct {
	Brand string
	Model string
}

func (c Car) Start() {
	fmt.Println("Car started")
}

func (c Car) Stop() {
	fmt.Println("Car stopped")
}

type Motorcycle struct {
	Brand string
	Model string
}

func (m Motorcycle) Start() {
	fmt.Println("Motorcycle started")
}

func (m Motorcycle) Stop() {
	fmt.Println("Motorcycle stopped")
}

func main() {
	car := Car{Brand: "Toyota", Model: "Camry"}
	motorcycle := Motorcycle{Brand: "Honda", Model: "CBR"}

	car.Start()
	car.Stop()

	motorcycle.Start()
	motorcycle.Stop()
}
