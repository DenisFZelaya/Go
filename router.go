package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary Muestra un mensaje
// @Description Obtiene un mensaje cuando se realiza un GET request
// @Tags ejemplos
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Este es un GET request"
// @Router /get [get]
func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Este es un GET request")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Este es un POST request")
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Este es un PUT request")
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Este es un DELETE request")
}

func main() {
	router := mux.NewRouter()
	r := gin.New()

	// EndPoints
	router.HandleFunc("/", getHandler).Methods("GET")
	router.HandleFunc("/", postHandler).Methods("POST")
	router.HandleFunc("/", putHandler).Methods("PUT")
	router.HandleFunc("/", deleteHandler).Methods("DELETE")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Println("Starting HTTP Server")
	http.ListenAndServe(":8080", router)

	r.Run()
}
