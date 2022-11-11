package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var saludazo = "Oh que paso mi niño"

var saludosCustom map[string]string

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

	if saludoCustom, ok := saludosCustom[params["idioma"]]; ok != false {
		saludo = saludoCustom
	}
	fmt.Fprintf(w, "%s\n", saludo)
}

func saludarPOST(w http.ResponseWriter, r *http.Request) {
	if saludosCustom == nil {
		saludosCustom = make(map[string]string)
	}

	defer r.Body.Close()
	var result struct {
		Idioma string `json:"idioma"`
		Saludo string `json:"saludo"`
	}
	json.NewDecoder(r.Body).Decode(&result)

	if result.Saludo != saludazo {
		saludazo = result.Saludo
	}
	fmt.Fprintf(w, "Creando nueva string para el idioma %s: %s\n", result.Idioma, result.Saludo)
	saludosCustom[result.Idioma] = result.Saludo
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/saludos/", saludarPOST).Methods("POST")
	r.HandleFunc("/saludos/{idioma:[a-zA-Z][a-zA-Z]}", saludarGET).Methods("GET")

	http.ListenAndServe(":8087", r)
}
