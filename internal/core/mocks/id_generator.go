package mocks

import (
	"github.com/google/uuid"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
)

//MockIDGenerator is mock of IDGenerator
type MockIDGenerator struct{}

//NewID mock
func (m *MockIDGenerator) NewID() domain.ID {
	return domain.ID(uuid.NewString())
}

//IDFromString mock
func (m *MockIDGenerator) IDFromString(str string) (domain.ID, error) {
	id, err := uuid.Parse(str)
	if err != nil {
		return domain.ZeroID, err
	}
	return domain.ID(id.String()), nil
}
