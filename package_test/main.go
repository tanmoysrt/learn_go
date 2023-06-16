package main

import (
	"fmt"
	"github.com/Tanmoy741127/test/utilshuh"
)

func main()  {
	fmt.Println("Hello World")
	fmt.Println(adder(1, 2))
	fmt.Println(utilshuh.Adder(1, 2))
	fmt.Println(utilshuh.Subber(1, 2))
	fmt.Println(utilshuh.Sum(1, 2))
}