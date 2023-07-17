package main

import "fmt"

func main() {
	// Condicional basico

	x := 10
	if x > 5 {
		fmt.Println("X is greater than 5")
	} else {
		fmt.Println("X is not greater than 5")
	}

	//Condicional con inicialización de variables (if con declaración)
	if y := 10; y > 5 {
		fmt.Println("y is greater than 5")
	}

	//  Condicional con operador NOT (!)
	x = 10

	if !(x > 5) {
		fmt.Println("x is not greater than 5")
	}

}
