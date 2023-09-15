package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pusupalahemanthkumar/pkg/routes"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server started at Port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
