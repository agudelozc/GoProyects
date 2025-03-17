package mocks

import "github.com/agudelozca/go-todo-api/models"

func (m *MockService) Register(username, password, role string) (*models.User, error) {
	args := m.Called(username, password, role)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockService) Authenticate(username, password string) (string, error) {
	args := m.Called(username, password)
	return args.String(0), args.Error(1)
}
