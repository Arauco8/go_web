package main

import (
	"fmt"
	"net/http"
)

func main() {

	server := http.NewServeMux() // Create a new ServeMux
	db :=
		server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status": "success", "message": "Welcome to the User API"}`))
		})
	server.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "success", "message": "User API"}`))
	})

	http.HandleFunc("/users", UserServer)

	fmt.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
