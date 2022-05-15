package handler

import (
	"log"
	"net/http"

	"FrontGate/pkg/mylogger"
)

var myLog *mylogger.MyLogger = mylogger.New()

// Home is home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	if err := RenderTemplate(w, "home_page.html"); err != nil {
		myLog.Info.Printf("Error parsing template: %s\n", err.Error())
	}
}

// About is about page handler
func About(w http.ResponseWriter, r *http.Request) {
	if err := RenderTemplate(w, "about_page.html"); err != nil {
		log.Fatalf("Error parsing template: %s\n", err.Error())
	}
}
