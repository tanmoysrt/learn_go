package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter rating : ")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)

	rating, err := strconv.ParseFloat(strings.TrimSpace(input), 64);
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Thanks for rating",rating+1)
	}




}