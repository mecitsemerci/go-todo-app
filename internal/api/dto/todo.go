package dto

import (
	"time"

	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
	"github.com/mecitsemerci/go-todo-app/internal/core/enum"
)

// TodoOutput is the result output dto of the todo item
type TodoOutput struct {
	ID          string             `json:"id" example:"5f68b3f08c111c96d1f8d9a3"`
	Title       string             `json:"title" example:"Shopping"`
	Description string             `json:"description" example:"Market shopping"`
	Completed   bool               `json:"completed" example:"false"`
	Priority    enum.PriorityLevel `json:"priority_level" example:"0"`
	CreatedAt   time.Time          `json:"created_at" example:"2020-07-28T07:32:32.71472Z"`
	UpdatedAt   time.Time          `json:"updated_at" example:"2020-07-30T07:32:32.71472Z"`
}

// FromModel is a mapping from Todo Model to TodoOutput
func (dto *TodoOutput) FromModel(todo todo.Todo) TodoOutput {
	dto.ID = todo.ID.String()
	dto.Title = todo.Title
	dto.Description = todo.Description
	dto.Completed = todo.Completed
	dto.Priority = todo.Priority
	dto.CreatedAt = todo.CreatedAt
	dto.UpdatedAt = todo.UpdatedAt
	return *dto
}

//CreateTodoInput is an input dto for a creation todo item
type CreateTodoInput struct {
	Title       string             `json:"title" binding:"required"`
	Description string             `json:"description,omitempty"`
	Priority    enum.PriorityLevel `json:"priority_level,omitempty"`
}

//ToModel is mapping from CreateTodoInput to Todo Model
func (dto *CreateTodoInput) ToModel() todo.Todo {
	return todo.Todo{
		Title:       dto.Title,
		Description: dto.Description,
		Priority:    dto.Priority,
	}
}

//CreateTodoOutput is the result output dto of the created Todo
type CreateTodoOutput struct {
	TodoID string `json:"todo_id"`
}

//UpdateTodoInput is an update input dto for given todo id
type UpdateTodoInput struct {
	Title       string             `json:"title,omitempty" binding:"required"`
	Description string             `json:"description,omitempty"`
	Priority    enum.PriorityLevel `json:"priority_level,omitempty"`
	Completed   bool               `json:"completed,omitempty"`
}

//ToModel is mapping from UpdateTodoInput to Todo Model
func (dto *UpdateTodoInput) ToModel(id string) todo.Todo {
	return todo.Todo{
		ID:          domain.ID(id),
		Title:       dto.Title,
		Priority:    dto.Priority,
		Description: dto.Description,
		Completed:   dto.Completed,
	}
}

//endregion
