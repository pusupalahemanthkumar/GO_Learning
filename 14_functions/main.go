package main

import (
	"fmt"
)

func main() {

	fmt.Println("functions in Golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Resilt is : ", result)

	fmt.Println("Pro Result is :", proAdder(1, 2, 3, 4, 5, 6, 7))

	rs, msg := twoValues()
	fmt.Println(rs, msg)

}

func adder(num1 int, num2 int) int {
	return num1 + num2

}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total

}

func twoValues() (int, string) {
	return 10, "Hello"

}

func greeter() {
	fmt.Println("Hello World !!")
}
