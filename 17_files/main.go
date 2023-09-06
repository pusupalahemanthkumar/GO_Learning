package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Golang")

	content := "This needs to go in a file !! !!!"

	file, err := os.Create("./demo.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("length is : ", length)

	defer file.Close()
	readFile("./demo.txt")

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNilError(err)

	fmt.Println("text data inside the file is \n", databyte)

	fmt.Println("text data inside the file is \n", string(databyte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}

}
