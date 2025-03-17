package tests

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/agudelozca/go-todo-api/internal/service"
	"github.com/agudelozca/go-todo-api/models"
	"github.com/agudelozca/go-todo-api/tests/mocks"
)

type TodoServiceTestSuite struct {
	suite.Suite
	service service.TodoService
	repo    *mocks.MockRepository
}

func (suite *TodoServiceTestSuite) SetupTest() {
	suite.repo = new(mocks.MockRepository)
	suite.service = service.NewTodoService(suite.repo)
}

func (suite *TodoServiceTestSuite) TestCreateTodo() {
	todo := models.Todo{
		Title: "test",
	}

	suite.repo.On("Create", mock.Anything).Return(nil)

	user, err := suite.service.CreateTodo(todo.Title)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), todo, user)
}

func (suite *TodoServiceTestSuite) TestGetTodo() {
	testUUID := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	todo := models.Todo{
		ID:    testUUID,
		Title: "test",
	}

	suite.repo.On("Get", mock.Anything).Return(todo, nil)

	result, err := suite.service.GetTodo(testUUID.String())
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), todo, result)
}
