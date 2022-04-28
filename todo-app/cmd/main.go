package main

import (
	"github.com/spf13/viper"
	"log"
	"todo"
	"todo/pkg/handler"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing config: %s\n", err)
	}

	handlers := handler.New()
	serv := new(todo.Server)

	if err := serv.Run(viper.GetString("port"), handlers.InitRouters()); err != nil {
		panic(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
