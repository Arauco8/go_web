package user

import (
	"context"
	"log"

	"github.com/Arauco8/go_web/internal/domain"
)

type (
	Service interface { // Interfaz que define los métodos para interactuar con el servicio de usuarios
		CreateUser(ctx context.Context, firstName, lastName, email string) (*domain.User, error) // Crea un nuevo usuario
		GetUserByID(ctx context.Context, id uint64) (*domain.User, error)                        // Obtiene un usuario por su ID
		GetAllUsers(ctx context.Context) ([]domain.User, error)                                  // Obtiene todos los usuarios
		UpdateUser(ctx context.Context, user *domain.User) error                                 // Actualiza un usuario existente
		DeleteUser(ctx context.Context, id uint64) error                                         // Elimina un usuario por su ID
		Close() error                                                                            // Cierra la conexión al servicio
	}

	service struct { // Estructura que implementa la interfaz Service
		repo Repository  // Repositorio para interactuar con la base de datos
		log  *log.Logger // Logger para registrar eventos
	}
)

func NewService(repo Repository, logger *log.Logger) Service { // Función que crea un nuevo servicio de usuarios
	return &service{
		repo: repo,
		log:  logger,
	}
}

func (s *service) CreateUser(ctx context.Context, firstName, lastName, email string) (*domain.User, error) { // Crea un nuevo usuario
	user := &domain.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
	err := s.repo.Create(ctx, user) // Llama al método Create del repositorio
	if err != nil {
		s.log.Printf("Error creating user: %v", err)
		return nil, err
	}
	s.log.Printf("User created: %v", user)
	return user, nil
}

func (s *service) GetUserByID(ctx context.Context, id uint64) (*domain.User, error) { // Obtiene un usuario por su ID
	user, err := s.repo.GetByID(ctx, id) // Llama al método GetByID del repositorio
	if err != nil {
		s.log.Printf("Error getting user by ID: %v", err)
		return nil, err
	}
	s.log.Printf("User found: %v", user)
	return user, nil
}

func (s *service) GetAllUsers(ctx context.Context) ([]domain.User, error) { // Obtiene todos los usuarios
	users, err := s.repo.GetAll(ctx) // Llama al método GetAll del repositorio
	if err != nil {
		s.log.Printf("Error getting all users: %v", err)
		return nil, err
	}
	s.log.Printf("All users found: %v", users)
	return users, nil
}

func (s *service) UpdateUser(ctx context.Context, user *domain.User) error { // Actualiza un usuario existente
	user, err := s.repo.GetByID(ctx, user.ID) // Llama al método GetByID del repositorio
	if err != nil {
		s.log.Printf("Error getting user by ID: %v", err)
		return err
	}
	err = s.repo.Update(ctx, user) // Llama al método Update del repositorio
	if err != nil {
		s.log.Printf("Error updating user: %v", err)
		return err
	}
	s.log.Printf("User updated: %v", user)
	return nil
}

func (s *service) DeleteUser(ctx context.Context, id uint64) error { // Elimina un usuario por su ID
	user, err := s.repo.GetByID(ctx, id) // Llama al método GetByID del repositorio
	if err != nil {
		s.log.Printf("Error getting user by ID: %v", err)
		return err
	}
	err = s.repo.Delete(ctx, user.ID) // Llama al método Delete del repositorio
	if err != nil {
		s.log.Printf("Error deleting user: %v", err)
		return err
	}
	s.log.Printf("User deleted: %v", user)
	return nil
}

func (s *service) Close() error { // Cierra la conexión al servicio
	return s.repo.Close() // Llama al método Close del repositorio
}
