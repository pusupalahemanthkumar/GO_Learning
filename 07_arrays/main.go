package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [5]string

	fruitList[0] = "apple"
	fruitList[1] = "banana"
	// fruitList[2] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit List is :", fruitList)
	fmt.Println("Fruit List Length is :", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy list is :", vegList)
	fmt.Println("vegy list length is :", len(vegList))

}
