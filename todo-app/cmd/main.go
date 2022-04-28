package main

import (
	"todo"
	"todo/pkg/handler"
)

func main() {
	handlers := handler.New()
	serv := new(todo.Server)

	if err := serv.Run("8080", handlers.InitRouters()); err != nil {
		panic(err)
	}
}
