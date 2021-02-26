package mocks

import (
	"context"

	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
)

//MockTodoRepository is mock TodoRepository
type MockTodoRepository struct {
	MockGetAll  func(ctx context.Context) ([]todo.Todo, error)
	MockGetByID func(ctx context.Context, id domain.ID) (todo.Todo, error)
	MockInsert  func(ctx context.Context, todo todo.Todo) (domain.ID, error)
	MockUpdate  func(ctx context.Context, todo todo.Todo) error
	MockDelete  func(ctx context.Context, id domain.ID) error
	MockClose   func(ctx context.Context) error
}

// GetAll mock
func (m *MockTodoRepository) GetAll(ctx context.Context) ([]todo.Todo, error) {
	return m.MockGetAll(ctx)
}

// GetByID mock
func (m *MockTodoRepository) GetByID(ctx context.Context, id domain.ID) (todo.Todo, error) {
	return m.MockGetByID(ctx, id)
}

// Insert mock
func (m *MockTodoRepository) Insert(ctx context.Context, todo todo.Todo) (domain.ID, error) {
	return m.MockInsert(ctx, todo)
}

// Update mock
func (m *MockTodoRepository) Update(ctx context.Context, todo todo.Todo) error {
	return m.MockUpdate(ctx, todo)
}

// Delete mock
func (m *MockTodoRepository) Delete(ctx context.Context, id domain.ID) error {
	return m.MockDelete(ctx, id)
}

// Close mock
func (m *MockTodoRepository) Close(ctx context.Context) error {
	return m.MockClose(ctx)
}
