package mypackage

import "fmt"

// CarPublic car con acceso publico
type CarPublic struct {
	Brand string
	Year int
}

type carPrivate struct {
	brand string
	year int
}

// PrintMessage imprimir un mensaje
func PrintMessage(name string) {
	fmt.Println("Hola", name)
}

func printMessage(name string) {
	fmt.Println("Hola", name)
}