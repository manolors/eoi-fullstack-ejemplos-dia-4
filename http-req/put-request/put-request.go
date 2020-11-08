package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//

func MakePutRequest() {

	mensaje := struct{ Data string }{"Hola mundo"}

	bytesMensaje, err := json.Marshal(mensaje)
	if err != nil {
		log.Fatalln(err)
	}
	const url = "https://httpbin.org/put"

	mensajePost := bytes.NewBuffer(bytesMensaje)
	req, err := http.NewRequest(http.MethodPut, url, mensajePost)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	for key, value := range result {
		fmt.Printf("%s: %v\n", key, value)
	}
}

func main() {
	fmt.Println("Put Request")
	MakePutRequest()
}
