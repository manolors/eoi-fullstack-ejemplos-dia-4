package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "ERROR")
		log.Println(r.Host, r.URL, r.RequestURI)
	})

	http.ListenAndServe(":8080", nil)
}
