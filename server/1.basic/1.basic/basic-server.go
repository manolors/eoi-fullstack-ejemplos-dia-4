package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola que tal\n")
		log.Println(r.Host, r.URL, r.RequestURI)
	})

	http.ListenAndServe(":8080", nil)

}
