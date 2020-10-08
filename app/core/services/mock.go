package services

import (
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"
)

type MockTodoService struct {
	MockGetAll func() ([]*todo.Todo, error)
	MockFind   func(id string) (*todo.Todo, error)
	MockCreate func(todo todo.Todo) (string, error)
	MockUpdate func(todo todo.Todo) error
	MockDelete func(id string) error
}

func (m *MockTodoService) GetAll() ([]*todo.Todo, error) {
	return m.MockGetAll()
}

func (m *MockTodoService) Find(id string) (*todo.Todo, error) {
	return m.MockFind(id)
}

func (m *MockTodoService) Create(todo todo.Todo) (string, error) {
	return m.MockCreate(todo)
}

func (m *MockTodoService) Update(todo todo.Todo) error {
	return m.MockUpdate(todo)
}

func (m *MockTodoService) Delete(id string) error {
	return m.MockDelete(id)
}
