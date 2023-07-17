package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>¡Hola desde el servidor Go!</title>
		</head>
		<body>
			<h1>¡Hola desde el servidor Go!</h1>
			<p>Este es un servidor HTTP simple que sirve archivos HTML.</p>
		</body>
		</html>
		`
		// Escribir html en la respuesta
		fmt.Fprintf(w, html)

	})

	// Iniciamos el servidor en el puerto 8080
	fmt.Println("Servidor iniciado en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
