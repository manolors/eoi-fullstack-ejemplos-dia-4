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

	// nota: esta url puede dejar de funcionar, acceder a https://webhook.site/ y obtener una nueva si da error
	const url = "https://webhook.site/4f49ff80-7525-43a4-9028-144dcd0de68e"
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
