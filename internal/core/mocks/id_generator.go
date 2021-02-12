package mocks

import (
	"github.com/google/uuid"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
)

type MockIDGenerator struct{}

func (m *MockIDGenerator) NewID() domain.ID {
	return domain.ID(uuid.NewString())
}

func (m *MockIDGenerator) IDFromString(str string) (domain.ID, error) {
	id, err := uuid.Parse(str)
	if err != nil {
		return domain.NilID, err
	}
	return domain.ID(id.String()), nil
}
