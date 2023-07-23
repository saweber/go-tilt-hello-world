package main

import (
	"fmt"
	"log"
	"net/http"
)

func baseRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", baseRoute)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	fmt.Println("application started")
	handleRequests()
}
