package todoDto

import (
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"
)

type CreateTodoInput struct {
	Title  *string `json:"title" binding:"required"`
}

func (dto *CreateTodoInput) ToEntity() todo.Todo {
	return todo.Todo{
		Title: *dto.Title,
	}
}
