package interfaces

import (
	"context"

	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
)

//TodoService represents use-cases
type TodoService interface {
	GetAll(ctx context.Context) ([]todo.Todo, error)
	Find(ctx context.Context, todoID domain.ID) (todo.Todo, error)
	Create(ctx context.Context, todo todo.Todo) (domain.ID, error)
	Update(ctx context.Context, todo todo.Todo) error
	Delete(ctx context.Context, todoID domain.ID) error
}
