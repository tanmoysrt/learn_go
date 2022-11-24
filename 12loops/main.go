package main

import "fmt"

func main()  {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for i := 0; i < len(days); i++ {
		println(days[i])
	}

	for _, day := range days {
		println(day)
	}

	tmp := 1

	for tmp < 10 {
		if tmp == 1{
			tmp++
			goto labelA
		}
		if tmp == 3 {
			tmp++
			continue
		}
		if tmp == 5 {
			break
		}
		println(tmp)
		tmp++
	}

	labelA:
	   fmt.Println(("Jumping at label A"))

}