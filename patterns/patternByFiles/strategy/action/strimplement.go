package action

import "fmt"

type RedCar struct{}

func (rc RedCar) Coloring() {
	fmt.Println("Coloring car in red")
}

type WhiteCar struct{}

func (wc WhiteCar) Coloring() {
	fmt.Println("Coloring car in white")
}

type BlackCar struct{}

func (bc BlackCar) Coloring() {
	fmt.Println("Coloring car in black")
}

type BlueCar struct{}

func (blc BlueCar) Coloring() {
	fmt.Println("Coloring car in blue")
}

type GreenCar struct{}

func (gc GreenCar) Coloring() {
	fmt.Println("Coloring car in green")
}
