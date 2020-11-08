package main

import (
	"fmt"
	"net/http"
	"sync"
)

func checkLink(link string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(link)
	if err != nil || resp.StatusCode >= 500 {
		fmt.Println(link, "puede que esté down")
	} else {
		fmt.Println(link, "está up")
	}
}

func main() {
	links := []string{
		"https://emezeta.com",
		"https://manolorodriguez.com",
		"https://twitter.com",
		"https://estonufunca.com",
		"https://httpbin.org/status/500",
	}

	wg := sync.WaitGroup{}
	for _, link := range links {
		wg.Add(1)
		go checkLink(link, &wg)
	}
	wg.Wait()
}
