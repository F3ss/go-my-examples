package main

import (
	"FrontGate/internal/handler"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	if err := http.ListenAndServe(portNumber, nil); err != nil {
		log.Fatalf("Error while started server, error: %s\n", err)
	}
}
