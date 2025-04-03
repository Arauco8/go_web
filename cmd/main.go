package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Arauco8/go_web/internal/domain"
	"github.com/Arauco8/go_web/internal/user"
)

func main() {
	server := http.NewServeMux() // Create a new ServeMux
	db := user.DB{
		Users: []domain.User{{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john_doe@gmail.com",
		}, {
			ID:        2,
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "jane_doe@gmail.com",
		}, {
			ID:        3,
			FirstName: "Alice",
			LastName:  "Smith",
			Email:     "alice_smith@gmail.com",
		}, {
			ID:        4,
			FirstName: "Bob",
			LastName:  "Johnson",
			Email:     "bob_johnson@gmail.com",
		}, {
			ID:        5,
			FirstName: "Charlie",
			LastName:  "Brown",
			Email:     "charlie_brown@gmail.com",
		}, {
			ID:        6,
			FirstName: "Diana",
			LastName:  "Prince",
			Email:     "diana_prince@gmail.com",
		}, {
			ID:        7,
			FirstName: "Eve",
			LastName:  "Adams",
			Email:     "eve_adams@gmail.com",
		}, {
			ID:        8,
			FirstName: "Frank",
			LastName:  "Miller",
			Email:     "frank_miller@gmail.com",
		}, {
			ID:        9,
			FirstName: "Grace",
			LastName:  "Hopper",
			Email:     "grace_hopper@gmail.com",
		}, {
			ID:        10,
			FirstName: "Hank",
			LastName:  "Pym",
			Email:     "hank_pym@gmail.com",
		}},
		MaxUserID: 3,
	}

	logger := log.New(os.Stdout, "user: ", log.Lshortfile|log.LstdFlags) //<-- Initialize logger

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
