package main

import "fmt"

func main() {
	fmt.Println("Methods in Golang")

	hemanth := User{"Hemanth", "hemanth@gamil.com", true, 23}

	fmt.Println(hemanth)

	hemanth.GetStatus()
	hemanth.NewMail("demo@go.dev")

	fmt.Println(hemanth)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active : ", u.Status)

}

func (u User) NewMail(e string) {
	u.Email = e
	fmt.Println("Email of this user is: ", u.Email)
}
