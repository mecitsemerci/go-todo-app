package todoDto

import "github.com/google/uuid"

type CreateTodoOutput struct {
	TodoId uuid.UUID `json:"todo_id"`
}
