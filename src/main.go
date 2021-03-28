package main

import "fmt"

type Car struct {
	brand string
	year int
}

func main() {
	
	mycar := Car{brand: "Ford", year: 2020}
	fmt.Println(mycar)

	// other way
	var otherCar Car
	otherCar.brand = "Ferrari"
	otherCar.year = 2020

	fmt.Println(otherCar)
}