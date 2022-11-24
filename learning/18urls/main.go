package main

import (
	"fmt"
	"net/url"
)

const myurl = "http://test.gorest.co.in:443/public/v2/users?_format=json&access-token=YOUR_ACCESS_TOKEN"

func main()  {
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	for k, _ := range result.Query() {
		fmt.Println(k, result.Query().Get(k))
	}
	fmt.Println(result.RawQuery)
}