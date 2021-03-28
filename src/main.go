package main

import c "fmt"

type pc struct {
	ram int
	disk int
	brand string
}

func (myPC pc) ping() {
	c.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRam() {
	myPC.ram = myPC.ram * 2
}

func main() {
	
	a := 50
	b := &a

	c.Println(a, b)
	c.Println(b, *b)

	*b = 100

	c.Println(a)

	myPC := pc{ram:16, disk: 512, brand: "apple"}
	c.Println(myPC)

	myPC.ping()

	c.Println(myPC)
	myPC.duplicateRam()
	c.Println(myPC)
	myPC.duplicateRam()
	c.Println(myPC)

}