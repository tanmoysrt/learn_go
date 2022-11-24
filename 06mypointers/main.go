package main

import "fmt"

func main()  {
	fmt.Println("fff");
	var ptr *int;

	myNumber := 26

	ptr = &myNumber

	fmt.Println("ptr val ", ptr);
	fmt.Println("ptr val ", *ptr);
	*ptr = 45;
	fmt.Println("ptr val ", *ptr);
}