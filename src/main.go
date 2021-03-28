package main

import (
	pk "curso-golang-platzi/src/mypackage"
	"fmt"
)

func main() {
	
	var mycar pk.CarPublic
	mycar.Brand = "Ferrari"
	mycar.Year = 2020

	fmt.Println(mycar)

	pk.PrintMessage("Fernando")

	

}