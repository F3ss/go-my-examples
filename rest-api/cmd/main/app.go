package main

import (
	"log"
	"net"
	"net/http"
	"rest_api/internal/user"
	"rest_api/pkg/logging"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logging.Init()
	log.Println("Create router")
	router := httprouter.New()

	userHandler := user.NewHandler()
	log.Println("Register user handler")
	userHandler.Register(router)

	if err := start(router); err != nil {
		panic(err)
	}
}

func start(router *httprouter.Router) error {
	log.Println("Start application")

	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		return err
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	return server.Serve(listener)
}
