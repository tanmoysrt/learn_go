package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Hello, World!")
}