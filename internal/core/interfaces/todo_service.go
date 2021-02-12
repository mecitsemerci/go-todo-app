package interfaces

import (
	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
)

type TodoService interface {
	GetAll() ([]todo.Todo, error)
	Find(todoId domain.ID) (*todo.Todo, error)
	Create(todo todo.Todo) (domain.ID, error)
	Update(todo todo.Todo) error
	Delete(todoId domain.ID) error
}
