package main

import "fmt"

func main() {
	
	m := make(map[string]int)

	m["jose"] = 14
	m["manuel"] = 20

	fmt.Println(m)

	// recorrer un map
	for i, valor := range m {
		fmt.Println(i, valor)
	}

	// encontrar un valor
	value, ok := m["josep"]
	fmt.Println(value, ok)

}