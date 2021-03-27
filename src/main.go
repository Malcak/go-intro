package main

import "fmt"

func main(){

	// declaracion de variables
	helloMessage := "Hello"
	worldMessage := "World"

	// print ln
	fmt.Println(helloMessage, worldMessage)

	// print f
	nombre := "platzi"
	cursos := 500
	
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	// sprint f

	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	// tipo de datos

	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

}