package main

import "fmt"

func main()  {
	num := 2

	switch num {
	case 1:
		fmt.Println("pressed  1")
	case 2:
		fmt.Println("pressed  2")
		fallthrough // to execute the next case
	case 3:
		fmt.Println("pressed  3")
		fallthrough
	default:
		fmt.Println("pressed  -1")
	}
	
}