package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var saludazo = "Oh que paso mi niño"

func saludarGET(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var saludo string
	switch params["idioma"] {
	case "es":
		saludo = "Que pasa tronco"
	case "en":
		saludo = "Hello There"
	case "cn":
		saludo = saludazo
	default:
		saludo = "👋"
	}
	fmt.Fprintf(w, "%s\n", saludo)
}

func saludarPOST(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	defer r.Body.Close()
	var result struct {
		Saludo string `json:"saludo"`
	}
	json.NewDecoder(r.Body).Decode(&result)

	if result.Saludo != saludazo {
		saludazo = result.Saludo
	}
	fmt.Fprintf(w, "Creando nueva string para el idioma %s: %s\n", params["idioma"], result.Saludo)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/saludos/{idioma:[a-zA-Z][a-zA-Z]}", saludarGET).Methods("GET")
	r.HandleFunc("/saludos/{idioma:[a-zA-Z][a-zA-Z]}", saludarPOST).Methods("POST")

	http.ListenAndServe(":8080", r)
}
