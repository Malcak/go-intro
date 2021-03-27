package main

import "fmt"

func main() {

	switch modulo := 4%2; modulo {
		case 0:
			fmt.Println("Es par")
		default:
			fmt.Println("Es impar")
	}

	// sin condicion especificada
	valor := 50
	switch {
		case valor > 100:
			fmt.Println("Es mayor que 100")
		case valor < 0:
			fmt.Println("Es menor que 0")
		default:
			fmt.Println("Sin condicion")
	}

}