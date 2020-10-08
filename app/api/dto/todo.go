package dto

import (
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/utils"
	"time"
)

type TodoOutput struct {
	Id          string    `json:"id" example:"5f68b3f08c111c96d1f8d9a3"`
	Title       string    `json:"title" example:"Shopping"`
	Description string    `json:"description" example:"Market shopping"`
	Completed   bool      `json:"completed" example:"false"`
	CreatedAt   time.Time `json:"created_at" example:"2020-07-28T07:32:32.71472Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2020-07-30T07:32:32.71472Z"`
}

func (dto *TodoOutput) FromEntity(todo todo.Todo) TodoOutput {
	return TodoOutput{
		Id:          todo.Id.Hex(),
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}

// region CREATE
type CreateTodoInput struct {
	Title       *string `json:"title" binding:"required"`
	Description *string `json:"description" binding:"required"`
}

func (dto *CreateTodoInput) ToEntity() todo.Todo {
	return todo.Todo{
		Title:       *dto.Title,
		Description: *dto.Description,
	}
}

type CreateTodoOutput struct {
	TodoId string `json:"todo_id"`
}
//endregion

//region UPDATE
type UpdateTodoInput struct {
	Title       *string `json:"title" binding:"required"`
	Description *string `json:"description" binding:"required"`
	Completed   *bool   `json:"completed,omitempty"`
}

func (dto *UpdateTodoInput) ToEntity(id string) (*todo.Todo, error) {
	oid, err := utils.OIDFromStr(id)
	if err != nil {
		return nil, err
	}
	return &todo.Todo{
		Id:          oid,
		Title:       *dto.Title,
		Description: *dto.Description,
		Completed:   *dto.Completed,
	}, nil

}
//endregion
