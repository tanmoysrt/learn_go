package main

import (
	"fmt"
	"net"
)

func main() {
	//establish connection
	connection, err := net.Dial("tcp", "0.0.0.0:3000")
	if err != nil {
			panic(err)
	}
	///send some data
	_, err = connection.Write([]byte("Hello Server! Greetings."))
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
			fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))
	defer connection.Close()
}