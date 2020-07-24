package dtos

import (
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"
	"time"
)

type TodoOutputDto struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (dto *TodoOutputDto) FromEntity(todo todo.Todo) TodoOutputDto {
	return TodoOutputDto{
		Id:        todo.Id.String(),
		Title:     todo.Title,
		IsDone:    todo.IsDone,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}
