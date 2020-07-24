package dtos

import "github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"

type CreateTodoInputDto struct {
	Title  *string `json:"title" binding:"required"`
}

func (dto *CreateTodoInputDto) ToEntity() todo.Todo {
	return todo.Todo{
		Title: *dto.Title,
	}
}
