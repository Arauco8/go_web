package user

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

// Este archivo contiene la definición de los controladores y las funciones relacionadas con el manejo de solicitudes HTTP para la entidad "Usuario".
// Los controladores son responsables de recibir las solicitudes HTTP, procesarlas y devolver las respuestas adecuadas al cliente.

type (
	Controller func(w http.ResponseWriter, r *http.Request) // Define una función que toma un http.ResponseWriter y un http.Request como argumentos
	// y no devuelve nada.
	// Esta función se puede usar como un controlador para manejar solicitudes HTTP.
	Endpoints struct { // Define una estructura que contiene varios controladores para manejar diferentes tipos de solicitudes HTTP para simplificar
		CreateUser  Controller // Controlador para crear un nuevo usuario
		GetUserByID Controller // Controlador para obtener un usuario por su ID
		GetAllUsers Controller // Controlador para obtener todos los usuarios
		UpdateUser  Controller // Controlador para actualizar un usuario existente
		DeleteUser  Controller // Controlador para eliminar un usuario por su ID
	}

	Createrequest struct { // Define una estructura que representa una solicitud para crear un nuevo usuario
		FirstName string `json:"first_name"` // Campo para el primer nombre del usuario
		LastName  string `json:"last_name"`  // Campo para el apellido del usuario
		Email     string `json:"email"`      // Campo para el correo electrónico del usuario
	}
)

func MakeEndpoints(ctx context.Context, s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		path := r.URL.Path

		switch r.Method {
		case http.MethodGet:
			if path == "/users" {
				GetAllUsers(ctx, s, w)
				return
			}
			if strings.HasPrefix(path, "/users/") {
				idStr := strings.TrimPrefix(path, "/users/")
				GetUserByID(ctx, s, w, idStr)
				return
			}
			handleNotFound(w)

		case http.MethodPost:
			if path == "/users" {
				var req Createrequest
				decode := json.NewDecoder(r.Body)
				if err := decode.Decode(&req); err != nil {
					w.WriteHeader(http.StatusBadRequest)
					response := Response{
						Status:  http.StatusBadRequest,
						Message: err.Error(),
					}
					json.NewEncoder(w).Encode(response)
					return
				}
				handlePost(ctx, s, w, req)
				return
			}
			handleNotFound(w)

		default:
			handleNotFound(w)
		}
	}
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // Usamos interface{} para Data
}

func GetAllUsers(ctx context.Context, s Service, w http.ResponseWriter) {
	users, err := s.GetAllUsers(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	DataResponse(w, http.StatusOK, users)
}

func GetUserByID(ctx context.Context, s Service, w http.ResponseWriter, idStr string) {
	// TODO: Implementar la lógica para convertir idStr a un tipo adecuado (int?)
	// y llamar al servicio para obtener el usuario por ID.
	// Por ahora, solo devolvemos un 404 simulado.
	handleNotFound(w)
}

func DataResponse(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	response := Response{
		Status:  status,
		Message: "success",
		Data:    data,
	}
	json.NewEncoder(w).Encode(response)
}

func handlePost(ctx context.Context, s Service, w http.ResponseWriter, req Createrequest) {
	if req.FirstName == "" || req.LastName == "" || req.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := Response{
			Status:  http.StatusBadRequest,
			Message: "missing field",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := s.CreateUser(ctx, req.FirstName, req.LastName, req.Email) // Llama al método CreateUser del servicio
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	DataResponse(w, http.StatusCreated, user) // Llama a la función DataResponse para enviar la respuesta
}

func handleNotFound(w http.ResponseWriter) {
	status := http.StatusNotFound
	w.WriteHeader(status)
	response := Response{
		Status:  status,
		Message: "not Found",
	}
	json.NewEncoder(w).Encode(response)
}
