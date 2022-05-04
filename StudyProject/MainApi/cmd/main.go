package main

import (
	mainapi "main_api"

	"github.com/sirupsen/logrus"
)

func main() {
	mainApiServer := mainapi.NewMainApiServer()

	if err := mainApiServer.RunMainApiServer("8080"); err != nil {
		logrus.Fatalf("Failed to start the server, error: %v\n", err)
	}

}
