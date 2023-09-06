package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang!!")

	// no inheritance in golang; No super or parent

	hemanth := User{"Hemanth", "hemanth@gmail.com", true, 23}
	fmt.Println(hemanth)
	fmt.Printf("Hemanth details are: %+v\n", hemanth)
	fmt.Println(hemanth.Name)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
