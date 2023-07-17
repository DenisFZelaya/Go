package main

import "fmt"

func main() {
	// Declaracion y asignacion
	var arr [5]int
	arr2 := [5]int{1, 2, 3, 4, 5}

	// Accesso y modificacion
	arr3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr3[0]) // Acceso al primer elemento (índice 0) del arreglo. Imprime: 1
	arr3[1] = 10         // Modificación del segundo elemento (índice 1) del arreglo.

	// Obtener longitud del arreglo
	arr4 := [5]int{1, 2, 3, 4, 5}
	length := len(arr4) // Obtiene la longitud del arreglo. En este caso, length será 5.

	// Iterar sobre un arreglo
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// Iterar con rango
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Arrays multidimensionales
	var matrix [3][3]int                                 // Declaración de un arreglo multidimensional de 3x3, inicializado con valores por defecto (0 en este caso).
	matrix := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}} // Declaración e inicialización de un arreglo multidimensional.

	// Punteros a arreglos
	arr := [5]int{1, 2, 3, 4, 5}
	ptr := &arr    // Obtiene el puntero al arreglo
	(*ptr)[0] = 10 // Modifica el primer elemento del arreglo a través del puntero.

}
