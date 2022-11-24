package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome := "Welcome to the playground!"	
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter any number : ")

	// comma ok || err ok
	input, err := reader.ReadString('\n')
	fmt.Println("Received ", input)
	fmt.Println(err)

}