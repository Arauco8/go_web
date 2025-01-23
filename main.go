package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func UserServer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGet(w)
	case http.MethodPost:
		handlePost(w)
	default:
		handleNotFound(w)
	}
}

func handleGet(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Status:  http.StatusOK,
		Message: "success in GET",
	}
	json.NewEncoder(w).Encode(response)
}

func handlePost(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Status:  http.StatusOK,
		Message: "success in POST",
	}
	json.NewEncoder(w).Encode(response)
}

func handleNotFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Status:  http.StatusNotFound,
		Message: "not Found",
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/users", UserServer)

	fmt.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
