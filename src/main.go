package main

import (
	"fmt"
)

func say (text string, c chan<- string) {
	c <- text
}

func printTextOurCh(c <-chan string) {
	fmt.Println(<-c)
}

func main() {
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", c)

	printTextOurCh(c)


}