package main

import (
	"log"
	"todoApp"
	handler "todoApp/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	server := new(todoApp.Server)
	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error with running app: %s", err.Error())
	}
}
