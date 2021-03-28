package main

import (
	"fmt"
	"strings"
)

func isPalindrome(str string) {
	var strReverse string
	str = strings.ToLower(str)
	str = strings.ReplaceAll(" ", str, str)
	for i := len(str) - 1; i >= 0; i-- {
		strReverse += string(str[i])
	}

	if str == strReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func main() {

	// 
	slice := []string{"Hola", "Que", "Hace"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	isPalindrome("Am a")

}