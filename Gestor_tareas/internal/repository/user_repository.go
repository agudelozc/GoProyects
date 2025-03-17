package repository

import (
	"errors"

	"github.com/agudelozca/go-todo-api/models"
)

type UserRepository interface {
	SaveUser(user *models.User) error
	FindUserByUsername(username string) (*models.User, error)
	FindUserByID(id string) (*models.User, error)
}

type InMemoryUserRepository struct {
	users map[string]*models.User
}

// NewInMemoryUserRepository crea una nueva instancia de InMemoryUserRepository.
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*models.User),
	}
}

// SaveUser guarda un usuario en memoria.
func (r *InMemoryUserRepository) SaveUser(user *models.User) error {
	if _, exists := r.users[user.Username]; exists {
		return errors.New("el usuario ya existe")
	}
	r.users[user.Username] = user
	return nil
}

// FindUserByUsername encuentra un usuario por su nombre de usuario.
func (r *InMemoryUserRepository) FindUserByUsername(username string) (*models.User, error) {
	user, exists := r.users[username]
	if !exists {
		return nil, nil
	}
	return user, nil
}

// FindUserByID encuentra un usuario por su ID.
func (r *InMemoryUserRepository) FindUserByID(id string) (*models.User, error) {
	for _, user := range r.users {
		if user.ID.String() == id {
			return user, nil
		}
	}
	return nil, nil
}
