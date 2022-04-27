package main

import "todo"

func main() {
	serv := new(todo.Server)
	if err := serv.Run("8080"); err != nil {
		panic(err)
	}
}
