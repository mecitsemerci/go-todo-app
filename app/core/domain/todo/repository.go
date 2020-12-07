package todo

import (
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain"
)

type Repository interface {
	GetAll() ([]*Todo, error)
	GetById(id domain.ID) (*Todo, error)
	Insert(todo Todo) (domain.ID, error)
	Update(todo Todo) error
	Delete(id domain.ID) error
}
