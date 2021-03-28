package main

import "fmt"

type polygon2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calculate(p polygon2D) {
	fmt.Println("Area:", p.area())
}

func main() {
	
	mySquare := cuadrado{base: 2}
	myRectangle := rectangulo{base: 2, altura: 4}

	calculate(mySquare)
	calculate(myRectangle)

	// lista de interfaces
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface)

}