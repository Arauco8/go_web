package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Arauco8/go_web/internal/user"
	"github.com/Arauco8/go_web/pkg/bootstrap"
)

func main() {
	server := http.NewServeMux() // Create a new ServeMux
	db := bootstrap.NewDB()      // Initialize the database

	logger := bootstrap.NewLogger() //<-- Initialize logger

	repo := user.NewRepository(db, logger)   // Create a new repository instance
	service := user.NewService(repo, logger) // Create a new service instance

	ctx := context.Background() // Para pasar el contexto a los controladores, la informacion a diferentes capas

	server.HandleFunc("/users", user.MakeEndpoints(ctx, service)) // Handle requests to /users

	port := ":8080"
	fmt.Println("Server listening on :8080")
	if err := http.ListenAndServe(port, server); err != nil { // Pasa 'server' aquí
		fmt.Println("Error starting server:", err)
	}
}
