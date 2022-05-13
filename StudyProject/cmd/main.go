package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page")
}

func AddValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	if err := http.ListenAndServe(portNumber, nil); err != nil {
		log.Fatalf("Error while started server, error: %s\n", err)
	}
}
