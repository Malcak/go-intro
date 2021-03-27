package main

import "fmt"

func main() {

	// defer

	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// continue y break

	for i := 0; i <10; i++ {

		if i == 2 {
			fmt.Printf("Es 2")
			continue
		}

		if i == 8 {
			fmt.Println("Es 8")
			break
		}

	}

}