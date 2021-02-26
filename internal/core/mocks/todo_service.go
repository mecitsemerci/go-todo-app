package mocks

import (
	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
)

// MockTodoService is mock of TodoService
type MockTodoService struct {
	MockGetAll func() ([]*todo.Todo, error)
	MockFind   func(id domain.ID) (*todo.Todo, error)
	MockCreate func(todo todo.Todo) (domain.ID, error)
	MockUpdate func(todo todo.Todo) error
	MockDelete func(id domain.ID) error
}

// GetAll mock
func (m *MockTodoService) GetAll() ([]*todo.Todo, error) {
	return m.MockGetAll()
}

// Find mock
func (m *MockTodoService) Find(id domain.ID) (*todo.Todo, error) {
	return m.MockFind(id)
}

// Create mock
func (m *MockTodoService) Create(todo todo.Todo) (domain.ID, error) {
	return m.MockCreate(todo)
}

// Update mock
func (m *MockTodoService) Update(todo todo.Todo) error {
	return m.MockUpdate(todo)
}

// Delete mock
func (m *MockTodoService) Delete(id domain.ID) error {
	return m.MockDelete(id)
}
