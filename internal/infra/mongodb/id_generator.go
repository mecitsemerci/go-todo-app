package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
)

// IDGenerator for mongodb
type IDGenerator struct{}

// NewIDGenerator returns mongodb IDGenerator
func NewIDGenerator() *IDGenerator {
	return &IDGenerator{}
}

// NewID returns domain ID for mongo objectID
func (i *IDGenerator) NewID() todo.ID {
	return todo.ID(primitive.NewObjectID().Hex())
}

// IDFromString converts string id to domain ID for mongo objectID
func (i *IDGenerator) IDFromString(str string) (todo.ID, error) {
	oid, err := primitive.ObjectIDFromHex(str)
	if err != nil {
		return todo.ZeroID, err
	}
	return todo.ID(oid.Hex()), err
}
