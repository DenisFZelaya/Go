package main

import "fmt"

func main() {
	// For basico
	for i := 0; i < 5; i++ {
		fmt.Println("i: ", i)
	}

	// Similar a while
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// Ciclo con range (para iterar sobre slices, arrays, mapas, etc.):
	arr := []int{1, 2, 3, 4, 5}
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
