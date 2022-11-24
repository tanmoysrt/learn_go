package main

import (
	"fmt"
	"sort"
)

func main()  {
	var fruitList = []string{};
	fruitList = append(fruitList, "apple");
	fruitList = append(fruitList, "apple2");
	fruitList = append(fruitList, "apple3");
	fruitList = append(fruitList, "apple4","apple5", "apple6");

	fruitList = append(fruitList[1:], "apple7");
	
	fmt.Printf("Type %T\n", fruitList)
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 456
	highScores[2] = 678
	highScores[3] = 899

	highScores = append(highScores, 555,666,321)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	// how to remove a value from slices based on index
	index := 2
	highScores = append(highScores[:index], highScores[index+1:]...)

	fmt.Println(highScores)

}