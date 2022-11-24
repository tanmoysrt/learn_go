package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

var endpoints = []string{"test"}

func main() {
	// go greeter("Hello")
	// greeter("World")

	websitelist := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
		"https://www.golang.org",
		"https://www.stackoverflow.com",
		"https://www.github.com",
		"https://www.youtube.com",
		"https://www.netflix.com",
		"https://www.twitter.com",
		"https://www.instagram.com",
		"https://www.linkedin.com",
		"https://www.reddit.com",
		"https://www.microsoft.com",
		"https://www.apple.com",
		"https://www.adobe.com",
		"https://www.spotify.com",
		"https://www.tiktok.com",
		"https://www.whatsapp.com",
		"https://www.tinder.com",
		"https://www.pinterest.com",
		"https://www.paypal.com",
	}

	for _, site := range websitelist {
		go getStatusCode(site)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(endpoints)
}

// func greeter(s string)  {
// 	  for i := 0; i < 6; i++ {
// 		fmt.Println(s)
// 	  }
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS something went wrong")
	}
	fmt.Printf("%d status code for %s\n", resp.StatusCode, endpoint)
	mut.Lock()
	endpoints = append(endpoints, endpoint)
	mut.Unlock()
}
