package todoDto

import (
	"github.com/google/uuid"
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"
)

type UpdateTodoInput struct {
	Title  *string `json:"title" binding:"required"`
	IsDone *bool   `json:"is_done,omitempty"`
}

func (dto *UpdateTodoInput) ToEntity(id uuid.UUID) todo.Todo {
	return todo.Todo{
		Id:     id,
		Title:  *dto.Title,
		IsDone: *dto.IsDone,
	}
}
