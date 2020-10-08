package todo

import (
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain"
)

type MockTodoRepository struct {
	MockGetAll  func() ([]*Todo, error)
	MockGetById func(id domain.ID) (*Todo, error)
	MockInsert  func(todo Todo) (domain.ID, error)
	MockUpdate  func(todo Todo) error
	MockDelete  func(id domain.ID) error
}

func (m *MockTodoRepository) GetAll() ([]*Todo, error) {
	return m.MockGetAll()
}

func (m *MockTodoRepository) GetById(id domain.ID) (*Todo, error) {
	return m.MockGetById(id)
}

func (m *MockTodoRepository) Insert(todo Todo) (domain.ID, error) {
	return m.MockInsert(todo)
}

func (m *MockTodoRepository) Update(todo Todo) error {
	return m.MockUpdate(todo)
}

func (m *MockTodoRepository) Delete(id domain.ID) error {
	return m.MockDelete(id)
}

