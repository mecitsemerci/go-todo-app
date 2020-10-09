package mongodb

import (
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Todo struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description"`
	Completed   bool               `bson:"completed"`
	CreatedAt   time.Time          `bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `bson:"updated_at,omitempty"`
}

func (e *Todo) FromModel(m *todo.Todo) error {
	oid, err := primitive.ObjectIDFromHex(m.Id.String())
	if err != nil {
		return err
	}
	e.Id = oid
	e.Title = m.Title
	e.Description = m.Description
	e.Completed = m.Completed
	e.CreatedAt = m.CreatedAt
	e.UpdatedAt = m.UpdatedAt
	return nil
}

func (e *Todo) ToModel() *todo.Todo {
	oid := ObjectId(e.Id)
	return &todo.Todo{
		Id:          &oid,
		Title:       e.Title,
		Description: e.Description,
		Completed:   e.Completed,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}
