// hello.go

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func MakeGetRequest() {
	const url = "https://manolorodriguez.com"
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var result map[string]interface{}

	for key, value := range result {
		fmt.Printf("%s: %v\n", key, value)
	}
}

func main() {
	MakeGetRequest()
}
