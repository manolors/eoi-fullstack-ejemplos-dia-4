package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func MakeGetRequest() {
	const url = "https://httpbin.org/status/500"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("status:", resp.Status)
	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	for key, value := range result {
		fmt.Printf("%s: %v\n", key, value)
	}
}

func MakePostRequest() {

	mensaje := struct{ Data string }{"Hola mundo"}

	bytesMensaje, err := json.Marshal(mensaje)
	if err != nil {
		log.Fatalln(err)
	}
	const url = "https://httpbin.org/status/404"
	const contentType = "application/json"
	mensajePost := bytes.NewBuffer(bytesMensaje)
	resp, err := http.Post(url, contentType, mensajePost)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("status:", resp.Status)

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	for key, value := range result {
		fmt.Printf("%s: %v\n", key, value)
	}
}

func main() {
	fmt.Println("Post Request")
	MakePostRequest()
	fmt.Println("Get Request")
	MakeGetRequest()
}
