package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//

func MakePostRequest() {

	mensaje := struct{ Data string }{"Hola mundo"}

	bytesMensaje, err := json.Marshal(mensaje)
	if err != nil {
		log.Fatalln(err)
	}
	const url = "https://webhook.site/528d2912-5a9d-4ebe-868f-e707f1f8c90f"
	const contentType = "application/json"
	mensajePost := bytes.NewBuffer(bytesMensaje)
	resp, err := http.Post(url, contentType, mensajePost)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	for key, value := range result {
		fmt.Printf("%s: %v\n", key, value)
	}
	resp.Body.Close()
}

func main() {
	fmt.Println("Post Request")
	MakePostRequest()
}
