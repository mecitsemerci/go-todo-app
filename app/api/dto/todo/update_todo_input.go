package todoDto

import (
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateTodoInput struct {
	Title       *string `json:"title" binding:"required"`
	Description *string `json:"description" binding:"required"`
	IsDone      *bool   `json:"is_done,omitempty"`
}

func (dto *UpdateTodoInput) ToEntity(id string) (todo.Todo, error) {

	if oid, err := primitive.ObjectIDFromHex(id); err != nil {
		return todo.Todo{}, err
	} else {
		return todo.Todo{
			Id:          oid,
			Title:       *dto.Title,
			Description: *dto.Description,
			IsDone:      *dto.IsDone,
		}, nil
	}

}
