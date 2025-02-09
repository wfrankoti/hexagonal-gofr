package main

import (
	"hexagonal-gofr/internal/core/services"
	"hexagonal-gofr/internal/infrastructure/http/handlers"
	"hexagonal-gofr/internal/infrastructure/repositories"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	// Inicializar el repositorio
	userRepo := repositories.NewUserRepositoryMemory()

	// Inicializar el servicio
	userService := services.NewUserService(userRepo)

	// Inicializar el manejador
	userHandler := handlers.NewUserHandler(userService)

	// Definir rutas
	app.POST("/users", userHandler.CreateUser)
	app.GET("/users/{id}", userHandler.GetUserByID)

	// Iniciar el servidor
	app.Run()
}
