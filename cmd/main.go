package main

import (
	"log"
	"todoApp"
	"todoApp/pkg/handler"
	"todoApp/pkg/repository"
	"todoApp/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todoApp.Server)
	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error with running app: %s", err.Error())
	}
}
