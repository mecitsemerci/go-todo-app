package redisdb

import (
	"github.com/google/uuid"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
)

//IDGenerator is mock of IDGenerator
type IDGenerator struct{}

//NewIDGenerator returns mongodb IDGenerator
func NewIDGenerator() *IDGenerator {
	return &IDGenerator{}
}

//NewID returns unique id
func (m *IDGenerator) NewID() domain.ID {
	return domain.ID(uuid.NewString())
}

//IDFromString
func (m *IDGenerator) IDFromString(str string) (domain.ID, error) {
	id, err := uuid.Parse(str)
	if err != nil {
		return domain.ZeroID, err
	}
	return domain.ID(id.String()), nil
}
