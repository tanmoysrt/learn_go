package main

import "fmt"

func main() {
	fmt.Println("Welcome");

	var fruitList [4]string;

	fruitList[0] = "a";
	fruitList[1] = "b";
	fruitList[2] = "c";
	fruitList[3] = "d";

	fmt.Println(fruitList)
	fmt.Println("Length ",len(fruitList))

	for i := 0; i < 4; i++ {
		fmt.Println(fruitList[i])
	}
}
