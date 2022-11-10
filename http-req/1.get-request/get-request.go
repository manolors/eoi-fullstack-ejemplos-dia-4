// hello.go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func MakeGetRequest() {
	const url = "https://httpbin.org/uuid"
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var result struct {
		Uuid string `json:"uuid"`
	}

	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Printf("%s\n", result.Uuid)
}

func main() {
	MakeGetRequest()
}
