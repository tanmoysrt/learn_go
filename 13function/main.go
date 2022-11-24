package main

import "fmt"

func main() {
	welcome("tanmoy")
	
	res, msg :=  divide(10, 11)
	fmt.Println(res, msg)
	adder(1,2,3,5,6)
}



func welcome(name string) {
	fmt.Println("Welcome ", name)
}

func divide(a int, b int)(int, string){
	return a/b, "Division is done"
}

func adder(values ...int){
	total := 0
	for _, value := range values {
		total += value
	}
	fmt.Println(total)
}