package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Arauco8/go_web/internal/domain"
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
	// Crea una nueva instancia de Endpoints y asigna los controladores a las funciones correspondientes
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetAllUsers(ctx, s, w)
		case http.MethodPost:
			w.Header().Set("Content-Type", "application/json")
			decode := json.NewDecoder(r.Body)
			var user domain.User
			if err := decode.Decode(&user); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				response := Response{
					Status:  http.StatusBadRequest,
					Message: err.Error(),
				}
				json.NewEncoder(w).Encode(response)
				return
			}
			handlePost(ctx, s, w, user)
		default:
			handleNotFound(w)
		}
	}
}

type Response struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []domain.User `json:"data,omitempty"`
}

func handleGet(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Status:  http.StatusOK,
		Message: "success in GET",
	}
	json.NewEncoder(w).Encode(response)
}

func GetAllUsers(ctx context.Context, s Service, w http.ResponseWriter) {
	users, error := s.GetAllUsers(ctx)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Response{
			Status:  http.StatusInternalServerError,
			Message: error.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	DataResponse(w, http.StatusOK, users)
}

func DataResponse(w http.ResponseWriter, status int, users interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	var userData []domain.User
	if u, ok := users.([]domain.User); ok {
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

func handlePost(context context.Context, s Service, w http.ResponseWriter, data interface{}) {
	req := data.(Createrequest) // Convierte el dato a tipo domain.User

	if req.FirstName == "" || req.LastName == "" || req.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := Response{
			Status:  http.StatusBadRequest,
			Message: "missing field",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	user, error := s.CreateUser(context, req.FirstName, req.LastName, req.Email) // Llama al método CreateUser del servicio
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := Response{
			Status:  http.StatusInternalServerError,
			Message: error.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	DataResponse(w, http.StatusCreated, user) // Llama a la función DataResponse para enviar la respuesta
	// Llama a la función DataResponse para enviar la respuesta
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
