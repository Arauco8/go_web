package user

import (
	"context"
	"log"
	"proyectos/GO_Web/internal/domain"
)

// Esquema de Base de datos
type DB struct {
	Users     []domain.User //lo obtenemos del package domain
	MaxUserID uint64
}

//Generamos Interfaz

type (
	Repository interface { // Interfaz que define los métodos para interactuar con la base de datos
		// CreateUser crea un nuevo usuario en la base de datos
		Create(ctx context.Context, user *domain.User) error
		// GetUserByID obtiene un usuario por su ID
		GetByID(ctx context.Context, id uint64) (*domain.User, error)
		// GetAllUsers obtiene todos los usuarios de la base de datos
		GetAll(ctx context.Context) ([]domain.User, error)
		// UpdateUser actualiza un usuario existente en la base de datos
		Update(ctx context.Context, user *domain.User) error
		// DeleteUser elimina un usuario de la base de datos
		Delete(ctx context.Context, id uint64) error
		// Close cierra la conexión a la base de datos
		Close() error
	}

	repo struct { /// Estructura que implementa la interfaz Repository
		db  DB
		log *log.Logger
	}
)

func NewRepository(db DB, logger *log.Logger) Repository { // Función que crea una nueva instancia de repo
	// y la inicializa con la base de datos y el logger proporcionados
	return &repo{
		db:  db,
		log: logger,
	}
}

func (r *repo) Create(ctx context.Context, user *domain.User) error {
	r.db.MaxUserID++
	user.ID = r.db.MaxUserID
	r.db.Users = append(r.db.Users, *user)
	r.log.Printf("User created: %v", user)
	return nil
}

func (r *repo) GetByID(ctx context.Context, id uint64) (*domain.User, error) {
	for _, user := range r.db.Users {
		if user.ID == id {
			r.log.Printf("User found: %v", user)
			return &user, nil
		}
	}

	return nil, domain.ErrUserNotFound{ID: id}
}

func (r *repo) GetAll(ctx context.Context) ([]domain.User, error) {
	r.log.Println("Repository get all users")
	return r.db.Users, nil
}

func (r *repo) Update(ctx context.Context, user *domain.User) error {
	for i, u := range r.db.Users {
		if u.ID == user.ID {
			r.log.Printf("User updated: %v", user)
			r.db.Users[i] = *user
			return nil
		}
	}
	return domain.ErrUserNotFound{ID: user.ID} // Devuelve una instancia del error
}

func (r *repo) Delete(ctx context.Context, id uint64) error {
	for i, user := range r.db.Users {
		if user.ID == id {
			r.log.Printf("User deleted: %v", user)
			r.db.Users = append(r.db.Users[:i], r.db.Users[i+1:]...)
			return nil
		}
	}
	return domain.ErrUserNotFound{ID: id} // Devuelve una instancia del error
}

func (r *repo) Close() error {
	// En este caso no hay una conexión real que cerrar, pero si hubiera, lo haríamos aquí.
	return nil
}
