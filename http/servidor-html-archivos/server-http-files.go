package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Manejador de ruta
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// usar http.FileServer para servir archivos
		http.ServeFile(w, r, "views/index.html")
	})

	http.HandleFunc("/contactos", func(w http.ResponseWriter, r *http.Request) {
		// usar http.FileServer para servir archivos
		http.ServeFile(w, r, "views/contactos.html")
	})

	// Iniciamos el servidor en el puerto 8080
	fmt.Println("Servidor iniciado en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
