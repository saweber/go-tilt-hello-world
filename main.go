package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func baseRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, greeting)
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", baseRoute)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

var greeting string

func main() {
	greeting = os.Getenv("GREETING")
	if greeting == "" {
		greeting = "?"
	}
	fmt.Println("application started")
	handleRequests()
}
