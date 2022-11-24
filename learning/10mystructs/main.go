package main

import "fmt"

func main() {
	fmt.Println()

	var tanmoy User = User{
		Name: "Tanmoy",
		Email: "ts754@sss.sss",
		Status: true,
		Age: 20,
	}
	fmt.Println(tanmoy)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
