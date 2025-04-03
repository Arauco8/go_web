package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/users", UserServer)

	fmt.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
