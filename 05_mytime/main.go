package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))

	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2000, time.April, 12, 9, 15, 0, 0, time.UTC)

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

}

/*

# TO SEE ENV
go env

# TO CREATE BUILD (executable file)
go build

# TO CREATE BUILD FOR OTHER MACHINES
GOOS="windows" go build
GOOS="linux" go build
GOOS="drawin" go build


*/
