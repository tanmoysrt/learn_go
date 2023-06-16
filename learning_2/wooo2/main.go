package main

import "fmt"

func main() {
	defer fmt.Println("Delayed !")
	defer fmt.Println("Delayed Last !")
	welcome("Tanmoy ")
	fmt.Println(adder(10, 20))
	fmt.Println(summ(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	tanmoy := User{"Tanmoy"}
	fmt.Println(tanmoy)
	tanmoy.SetName("Tanmoy Sarkar")
	fmt.Println(tanmoy)
}

func welcome(name string) {
	fmt.Println("Welcome to "+name)
}

func adder(x, y int) int {
	return x + y
}

func summ(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// Pointer to update old value
func (uptr *User) SetName(name string) {
	uptr.FirstName = name
}

type User struct {
	FirstName string
}