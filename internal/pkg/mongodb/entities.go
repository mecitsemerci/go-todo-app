package mongodb

import (
	"time"

	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
	"github.com/mecitsemerci/go-todo-app/internal/core/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Todo entity of mongodb
type Todo struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Completed   bool               `bson:"completed"`
	Priority    uint8              `bson:"priority_level"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

//FromModel mapper from todo model to todo entity
func (e *Todo) FromModel(m *todo.Todo) error {
	oid, err := primitive.ObjectIDFromHex(string(m.ID))
	if err != nil {
		return err
	}
	e.ID = oid
	e.Title = m.Title
	e.Description = m.Description
	e.Completed = m.Completed
	e.Priority = uint8(m.Priority)
	e.CreatedAt = m.CreatedAt
	e.UpdatedAt = m.UpdatedAt
	return nil
}

//ToModel mapper from todo entity to todo model
func (e *Todo) ToModel() todo.Todo {
	return todo.Todo{
		ID:          domain.ID(e.ID.Hex()),
		Title:       e.Title,
		Description: e.Description,
		Completed:   e.Completed,
		Priority:    enum.PriorityLevel(e.Priority),
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}
