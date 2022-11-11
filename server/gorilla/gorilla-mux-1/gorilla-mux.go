package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func persona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "%s %s\n", vars["nombre"], vars["apellidos"])
}

func suma(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	first, _ := strconv.Atoi(vars["first"])
	second, _ := strconv.Atoi(vars["second"])

	fmt.Fprintf(w, "%d\n", first+second)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/persona/{nombre}/{apellidos}/", persona)
	r.HandleFunc("/suma/{first}/{second}/", suma).Methods("POST")

	http.ListenAndServe(":8049", r)
}
