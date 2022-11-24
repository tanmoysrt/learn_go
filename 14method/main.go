package main

import "fmt"

func main() {
	var tanmoy User = User{
		Name: "Tanmoy",
		Email: "ts754@sss.sss",
		Status: true,
		Age: 20,
	}
	tanmoy.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}


func (u User) GetStatus() {
	if u.Status {
		fmt.Println("Active")
	} else {
		fmt.Println("Inactive")
	}
}