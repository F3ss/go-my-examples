package main

import (
	mainapi "main_api"
	"main_api/internal/handler"

	"github.com/sirupsen/logrus"
)

func main() {
	mainApiServer := mainapi.NewMainApiServer()
	handler := handler.NewHandler()

	if err := mainApiServer.RunMainApiServer("8080", handler.InitRoutes()); err != nil {
		logrus.Fatalf("Failed to start the server, error: %v\n", err)
	}

}
