// hello.go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const url = "https://httpbin.org/delay/"

func MakeGetRequestCustom(delay string) error {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url + delay)
	if err != nil {
		return err
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	for key, value := range result {
		fmt.Printf("%s: %v\n", key, value)
	}
	return nil
}

func main() {
	fmt.Println("Esta request debería petar:")
	err := MakeGetRequestCustom("6")
	if err != nil {
		fmt.Println("Error en la request:")
		fmt.Println(err)
	}

	fmt.Println("\n\nEsta request debería funcionar:")
	err = MakeGetRequestCustom("4")
	if err != nil {
		fmt.Println("Error en la request:")
		fmt.Println(err)
	}
}
