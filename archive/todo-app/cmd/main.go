package main

import (
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repositories"
	"todo/pkg/service"
)

func main() {
	repos := repositories.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	serv := new(todo.Server)

	if err := serv.Run("8080", handlers.InitRouters()); err != nil {
		panic(err)
	}
}
