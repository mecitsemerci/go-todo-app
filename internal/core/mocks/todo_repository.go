package mocks

import (
	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
)

type MockTodoRepository struct {
	MockGetAll  func() ([]todo.Todo, error)
	MockGetById func(id domain.ID) (*todo.Todo, error)
	MockInsert  func(todo todo.Todo) (domain.ID, error)
	MockUpdate  func(todo todo.Todo) error
	MockDelete  func(id domain.ID) error
}

func (m *MockTodoRepository) GetAll() ([]todo.Todo, error) {
	return m.MockGetAll()
}

func (m *MockTodoRepository) GetById(id domain.ID) (*todo.Todo, error) {
	return m.MockGetById(id)
}

func (m *MockTodoRepository) Insert(todo todo.Todo) (domain.ID, error) {
	return m.MockInsert(todo)
}

func (m *MockTodoRepository) Update(todo todo.Todo) error {
	return m.MockUpdate(todo)
}

func (m *MockTodoRepository) Delete(id domain.ID) error {
	return m.MockDelete(id)
}

