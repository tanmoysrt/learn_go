package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	var username string = "tanmoy"
	fmt.Printf("Hello, %s\n", username)
	fmt.Printf("Type: %T\n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)

	var age int = 25
	fmt.Println(age)

	var height float64 = 5.6
	fmt.Println(height)

	{
		var height float32 = 5.99
		fmt.Println(height)
	}

	number := 10
	fmt.Println(number)

	const pi float64 = 3.14159265359
	fmt.Println(pi)

	ptr := &number
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = 20
	fmt.Println(number)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter rating: ")
	text, _ := reader.ReadString('\n')
	fmt.Print(text)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(text), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(numRating)
	}

	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05 Mon January")
	fmt.Println(formattedTime)
	fmt.Println(currentTime.UnixMilli())

	// var arr = []int{1, 2, 3}
	arr := make([]int, 3)
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)
	fmt.Println(len(arr))

	arr = append(arr, 4, 5)
	fmt.Println(arr)

	arr = append(arr[1:], )
	fmt.Println(arr)

	arr = append(arr[:len(arr)-1])
	fmt.Println(arr)

	// Insert
	arr = append(arr[:1], append([]int{100}, arr[1:]...)...)
	fmt.Println(arr)

	fmt.Println(sort.IntsAreSorted(arr))
	sort.Ints(arr)
	fmt.Println(arr)
	fmt.Println(sort.IntsAreSorted(arr))

	languages := make(map[string]string)
	languages["js"] = "JavaScript"
	languages["py"] = "Python"

	fmt.Println(languages)
	delete(languages, "py")
	fmt.Println(languages)
	fmt.Println(languages["js=="])

	for key, value := range languages {
		fmt.Println(key, value)
	}

	tanmoy := User{"Tanmoy", "SArkar"}
	fmt.Println(tanmoy)

	if tanmoy.FirstName == "Tanmoy" {
		fmt.Println("Tanmoy")
	} else {
		fmt.Println("Not Tanmoy")
	}

	diceNumber := 1
	switch diceNumber {
	case 1, 3, 5:
		fmt.Println("Odd")
		// it will fallthrough to the next case
		// fallthrough
	case 2, 4, 6:
		fmt.Println("Even")
	}

	for i := 0; i < 5; i++ {
		if i == 1 {
			continue
		}
		fmt.Println(i)

		if i == 3 {
			break
		}
	}


}


type User struct {
	FirstName string
	LastName  string
}