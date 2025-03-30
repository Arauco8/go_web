package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

var users []User
var maxID uint64

func init() {
	users = []User{{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john_doe@hotmail.com",
	},
		{
			ID:        2,
			FirstName: "Jane",
			LastName:  "Smith",
			Email:     "jane_smith@gmail.com",
		},
		{
			ID:        3,
			FirstName: "Alice",
			LastName:  "Johnson",
			Email:     "alice_johnson",
		}}
	maxID = 3
	// Initialize the users slice with some sample data
	// You can add more users here if needed
	// For example:
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data,omitempty"`
}

func UserServer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetAllUsers(w)
	case http.MethodPost:
		handlePost(w, r)
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

func GetAllUsers(w http.ResponseWriter) {
	DataResponse(w, http.StatusOK, users)
}

func DataResponse(w http.ResponseWriter, status int, users interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	var userData []User
	if u, ok := users.([]User); ok {
		userData = u
	}

	response := Response{
		Status:  status,
		Message: "success",
		Data:    userData,
	}
	json.NewEncoder(w).Encode(response) // Encode the response

	/*Usa json.Marshal cuando necesites la representación JSON como un valor en tu código.
	     Usa json.NewEncoder(w).Encode cuando necesites enviar la respuesta JSON directamente
		 a un flujo de salida, como una respuesta HTTP o un archivo, especialmente si la estructura
		 de datos puede ser grande o si necesitas un flujo de escritura continuo.
	*/
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decode := json.NewDecoder(r.Body)
	var user User
	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := Response{
			Status:  http.StatusBadRequest,
			Message: "missing field",
		}
		json.NewEncoder(w).Encode(response)
		return

	}

	if err := decode.Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	response := Response{
		Status:  http.StatusCreated,
		Message: "success in POST",
		Data:    []User{user},
	}
	maxID++
	user.ID = maxID // Assign a new ID to the user
	users = append(users, user)
	json.NewEncoder(w).Encode(response)
}

func handleNotFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	status := http.StatusNotFound
	w.WriteHeader(status)
	response := Response{
		Status:  status,
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
