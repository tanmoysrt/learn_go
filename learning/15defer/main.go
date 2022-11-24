package main

import "fmt"

func main()  {
	defer fmt.Println("Hello 1");
	defer fmt.Println("Hello 2");
	fmt.Println("World");
}