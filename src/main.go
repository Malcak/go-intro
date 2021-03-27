package main

import "fmt"

func helloWorld(x string) {
	fmt.Println("Hola mundo ", x)
}

func tripeArgument(a int, b int, c string) {
 fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a*2
}

func main() {

	helloWorld("2")
	tripeArgument(1, 2, "hola")
	fmt.Println(returnValue(5))
	value1, value2 := doubleReturn(5)
	fmt.Println(value1, value2)

	value3, _ := doubleReturn(6)
	fmt.Println(value3)

}