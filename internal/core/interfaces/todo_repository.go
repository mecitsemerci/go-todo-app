package interfaces

import (
	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
)

type TodoRepository interface {
	GetAll() ([]todo.Todo, error)
	GetById(id domain.ID) (*todo.Todo, error)
	Insert(todo todo.Todo) (domain.ID, error)
	Update(todo todo.Todo) error
	Delete(id domain.ID) error
}
