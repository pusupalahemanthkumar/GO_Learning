package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco/dev:3000/learn?coursename=reactjs&&paymentid=vghfhfg"

func main() {
	fmt.Println("Welcome to handling URLS in Golang")
	// fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparam := result.Query()
	fmt.Printf("type of query params are : %T\n", qparam)
	fmt.Println(qparam["coursename"])

	for _, val := range qparam {
		fmt.Println("Param is : ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "loc.dev",
		Path:     "/tutcss",
		RawQuery: "user=hemanth",
	}

	anootherUrl := partsOfUrl.String()

	fmt.Println(anootherUrl)

}
