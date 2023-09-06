package main

import "fmt"

func main() {
	fmt.Println("loops in Golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "saturday"}

	fmt.Println(days)

	// Method-1
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	//Method-2
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	//Method-3
	for index, day := range days {
		fmt.Println(index, day)
	}

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 2 {
			rougeValue++
			goto lco
		}
		if rougeValue == 3 {
			rougeValue++
			continue
		}
		if rougeValue == 5 {
			break
		}
		fmt.Println("value is : ", rougeValue)
		rougeValue++

	lco:
		fmt.Println("Jumping !!")
	}

}
