package todoDto

import (
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"
	"time"
)

type TodoOutput struct {
	Id        string    `json:"id" example:"6ba7b811-9dad-11d1-80b4-00c04fd430c8"`
	Title     string    `json:"title" example:"Shopping"`
	IsDone    bool      `json:"is_done" example:"false"`
	CreatedAt time.Time `json:"created_at" example:"2020-07-28T07:32:32.71472Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-07-30T07:32:32.71472Z"`
}

func (dto *TodoOutput) FromEntity(todo todo.Todo) TodoOutput {
	return TodoOutput{
		Id:        todo.Id.String(),
		Title:     todo.Title,
		IsDone:    todo.IsDone,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}
