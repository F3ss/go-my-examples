package main

import (
	"log"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	parseTemplate, err := template.ParseFiles("./home_page.html")
	if err != nil {
		log.Fatalf("Error parsing template: %s\n", err.Error())
	}

	parseTemplate.Execute(w, nil)
}

func About(w http.ResponseWriter, r *http.Request) {
	parseTemplate, err := template.ParseFiles("./about_page.html")
	if err != nil {
		log.Fatalf("Error parsing template: %s\n", err.Error())
	}

	parseTemplate.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	if err := http.ListenAndServe(portNumber, nil); err != nil {
		log.Fatalf("Error while started server, error: %s\n", err)
	}
}
