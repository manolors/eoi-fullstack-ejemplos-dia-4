package main

import (
	"fmt"
	"net/http"
)

func saludoIngles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello There\n")
}

func saludoCanario(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Oh que pasó mi niño\n")
}

func saludoEspañol(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Que pasa tio\n")
}

func main() {
	http.HandleFunc("/ingles", saludoIngles)
	http.HandleFunc("/español", saludoEspañol)
	http.HandleFunc("/canario/", saludoCanario)
	http.ListenAndServe(":8090", nil)
}
