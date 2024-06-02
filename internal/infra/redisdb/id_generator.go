package redisdb

import (
	"github.com/google/uuid"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
)

// IDGenerator is mock of IDGenerator
type IDGenerator struct{}

// NewIDGenerator returns mongodb IDGenerator
func NewIDGenerator() *IDGenerator {
	return &IDGenerator{}
}

// NewID returns unique id
func (m *IDGenerator) NewID() todo.ID {
	return todo.ID(uuid.NewString())
}

// IDFromString
func (m *IDGenerator) IDFromString(str string) (todo.ID, error) {
	id, err := uuid.Parse(str)
	if err != nil {
		return todo.ZeroID, err
	}
	return todo.ID(id.String()), nil
}
