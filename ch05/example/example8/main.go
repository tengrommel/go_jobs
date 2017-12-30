package main

import "fmt"

type Car struct {
	weight int
	name string
}

func (p *Car) Run() {
	fmt.Println("running")
}

type Bike struct {
	Car
	lunzi int
}

func main() {
	var a Bike
	a.weight = 100
	a.name = "bike"
	a.lunzi = 2
	fmt.Println(a)
	a.Run()
}
