package interfaces

import (
	"context"

	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
)

// TodoRepository interface guides all operations
type TodoRepository interface {
	GetAll(ctx context.Context) ([]todo.Todo, error)
	GetByID(ctx context.Context, id domain.ID) (todo.Todo, error)
	Insert(ctx context.Context, todo todo.Todo) (domain.ID, error)
	Update(ctx context.Context, todo todo.Todo) error
	Delete(ctx context.Context, id domain.ID) error
	Close(ctx context.Context) error
}
