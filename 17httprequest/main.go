package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://gorest.co.in/public/v2/users"

func main()  {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response type -> %T\n", response)
	
	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(databytes)
	fmt.Println(string(databytes))

}