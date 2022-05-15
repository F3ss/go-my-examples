package main

import (
	"log"
	"net/http"
	"text/template"

	"FrontGate/pkg/mylogger"
)

const portNumber = ":8080"

var myLog *mylogger.MyLogger = mylogger.New()

func home(w http.ResponseWriter, r *http.Request) {
	if err := renderTemplate(w, "home_page.html"); err != nil {
		myLog.Info.Printf("Error parsing template: %s\n", err.Error())
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	if err := renderTemplate(w, "about_page.html"); err != nil {
		log.Fatalf("Error parsing template: %s\n", err.Error())
	}
}

func renderTemplate(w http.ResponseWriter, templateName string) error {
	parseTemplate, err := template.ParseFiles("./templates/" + templateName)
	if err != nil {
		return err
	}

	if err := parseTemplate.Execute(w, nil); err != nil {
		return err
	}

	return nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	if err := http.ListenAndServe(portNumber, nil); err != nil {
		log.Fatalf("Error while started server, error: %s\n", err)
	}
}
