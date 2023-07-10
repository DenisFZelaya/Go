package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// Struct
type DataAdd struct {
	N1    int `json:"n1"`
	N2    int `json:"n2"`
	Total int `json:"total"`
}

// EndPoint Inicial
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func myNameHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener parametros
	params := r.URL.Query()
	// Obtener el parametro
	name := params.Get("name")
	// Response
	fmt.Fprintln(w, "!Hola, ", name, "!")
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	n1, err1 := strconv.Atoi(params.Get("n1"))

	if err1 != nil {
		fmt.Println("ERROR:", err1)
	}

	n2, err2 := strconv.Atoi(params.Get("n2"))

	if err2 != nil {
		fmt.Println("ERROR:", err2)
	}
	total := n1 + n2

	fmt.Fprintf(w, "La suma de: %d + %d = %d", n1, n2, total)
}

func addJSONHandler(w http.ResponseWriter, r *http.Request) {

	data := DataAdd{
		N1:    6,
		N2:    8,
		Total: 14,
	}

	// Establecer el contenido de respuesta JSON
	w.Header().Set("Content-Type", "application/json")

	// Codifica a JSON y responde
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func addJSONParamsHandler(w http.ResponseWriter, r *http.Request) {

	data := DataAdd{}

	params := r.URL.Query()

	n1, err1 := strconv.Atoi(params.Get("n1"))

	if err1 != nil {
		fmt.Println("ERROR:", err1)
	}

	n2, err2 := strconv.Atoi(params.Get("n2"))

	if err2 != nil {
		fmt.Println("ERROR:", err2)
	}
	total := n1 + n2

	data.N1 = n1
	data.N2 = n2
	data.Total = total

	// Establecer el contenido de respuesta JSON
	w.Header().Set("Content-Type", "application/json")

	// Codifica a JSON y responde
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func readBodyHandler(w http.ResponseWriter, r *http.Request) {

	_, err := io.Copy(w, r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func readBodyDataHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	name := data["name"].(string)
	lastname := data["lastname"].(string)

	w.Write([]byte("Nombre: " + name + ", Apellido: " + lastname))
}

type Person struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	FullName string `json:"fullname"`
}

func readBodyDataStructHandler(w http.ResponseWriter, r *http.Request) {
	var p Person
	err1 := json.NewDecoder(r.Body).Decode(&p)

	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	p.FullName = "Nombre completo: " + p.Name + " " + p.LastName

	// Codifica a JSON y responde
	err := json.NewEncoder(w).Encode(p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func crudHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "Este es un GET request")
	case "POST":
		fmt.Fprintln(w, "Este es un POST request")
	case "PUT":
		fmt.Fprintln(w, "Este es un PUT request")
	case "DELETE":
		fmt.Fprintln(w, "Este es un DELETE request")
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func server() {
	http.HandleFunc("/", crudHandler)
	http.HandleFunc("/name", myNameHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/addJSON", addJSONHandler)
	http.HandleFunc("/addJSONParams", addJSONParamsHandler)
	http.HandleFunc("/body", readBodyHandler)
	http.HandleFunc("/bodyJSON", readBodyDataHandler)
	http.HandleFunc("/bodyJSONtoStruct", readBodyDataStructHandler)

	fmt.Println("Starting HTTP Server")
	http.ListenAndServe(":8080", nil)
}
