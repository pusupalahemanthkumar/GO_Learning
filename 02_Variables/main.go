package main

import "fmt"

// Below code is not allowed (It is allowed inside method)
// jwtToken := 3000

// Capital Letter First ==> Public
const LoginToken string = "GZSVHBBCVH"

func main() {

	// String Type

	var username string = "Hemanth"
	fmt.Println(username)
	fmt.Printf("Varible is of type : %T \n", username)

	// Bool Type

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Varible is of type : %T \n", isLoggedIn)

	// Integer Types

	// var smallVal uint8 = 255
	// var smallVal uint8 = 256
	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("Varible is of type : %T \n", smallVal)

	// Float Types

	// var smallFloat float32 = 256.7788997666
	var smallFloat float64 = 256.7788997666
	fmt.Println(smallFloat)
	fmt.Printf("Varible is of type : %T \n", smallFloat)

	// Default Values

	var anotherVariable int
	fmt.Println(anotherVariable) // 0
	fmt.Printf("Varible is of type : %T \n", anotherVariable)

	// Implicit Type

	var website = "https://go.dev/dl/"
	fmt.Println(website)
	fmt.Printf("Varible is of type : %T \n", website)

	// No var style (It is allowed inside method)

	numberOfUser := 300000
	fmt.Println(numberOfUser)
	fmt.Printf("Varible is of type : %T \n", numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Varible is of type : %T \n", LoginToken)

}
