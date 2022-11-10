// hello.go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const url = "https://httpbin.org/get"

func MakeGetRequest() {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	for key, value := range result {
		fmt.Printf("%s: %v\n", key, value)
	}
}

func MakeGetRequestCustom() {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	for key, value := range result {
		fmt.Printf("%s: %v\n", key, value)
	}
}

func main() {
	MakeGetRequestCustom()
}
