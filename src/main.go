package main

import "fmt"

type pc struct {
	ram int
	disk int
	brand string
}

func (myPC pc) String() string {
	return fmt.Sprintf("{ram:%d, disk:%d, brand:%s}", myPC.ram, myPC.disk, myPC.brand)
}

func main() {
	
	myPC := pc{ram: 16, disk: 512, brand: "apple"}

	fmt.Println(myPC)

}