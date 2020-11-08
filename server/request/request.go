package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var result map[string]interface{}
	json.NewDecoder(r.Body).Decode(&result)
	for key, value := range result {
		fmt.Printf("%s: %v\n", key, value)
	}
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}
