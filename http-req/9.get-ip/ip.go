package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetIp() {
	const url = "https://httpbin.org/ip"
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode == http.StatusOK {
		var result struct {
			Origin string `json:"origin"`
		}
		json.NewDecoder(resp.Body).Decode(&result)
		fmt.Printf("%s", result.Origin)
	}

	if resp.StatusCode == http.StatusInternalServerError {
		log.Println("El servidor respondió un error")
	}

}

func main() {
	fmt.Println("Get uuid")
	GetIp()
}
