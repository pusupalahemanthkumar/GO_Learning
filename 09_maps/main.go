package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in Golang !!")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	fmt.Println("List of all languages : ", languages)
	fmt.Println("JS short for : ", languages["JS"])
	fmt.Println("JA short for : ", languages["JA"])

	delete(languages, "RB")
	fmt.Println("List of all languages : ", languages)

	// loops
	for key, value := range languages {
		fmt.Printf("For Key %v the value is : %v\n", key, value)
	}

}
