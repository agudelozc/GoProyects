// internal/service/user_service.go
package service

import (
	"errors"
	"log"
	"time"

	"github.com/agudelozca/go-todo-api/internal/repository"
	"github.com/agudelozca/go-todo-api/models"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(username, password, role string) (*models.User, error)
	Authenticate(username, password string) (string, error) // Retorna token JWT
}

const tokenExpiryHours = 72

type userService struct {
	repo      repository.UserRepository
	jwtSecret []byte
}

func NewUserService(repo repository.UserRepository, secret []byte) UserService {
	return &userService{
		repo:      repo,
		jwtSecret: secret,
	}
}

func (s *userService) Register(username, password, role string) (*models.User, error) {
	if username == "" || password == "" {
		return nil, errors.New("el nombre de usuario y la contraseña son obligatorios")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	log.Printf("Contraseña hasheada: %s", hashedPassword)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:       uuid.New(),
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
	}

	err = s.repo.SaveUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Authenticate(username, password string) (string, error) {
	user, err := s.repo.FindUserByUsername(username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("usuario no encontrado")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("contraseña inválida")
	}

	// Generar token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID.String(),
		"exp":      time.Now().Add(time.Hour * tokenExpiryHours).Unix(),
		"role":     user.Role,
		"username": user.Username,
	})

	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
