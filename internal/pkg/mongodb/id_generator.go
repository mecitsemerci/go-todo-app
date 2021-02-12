package mongodb

import (
	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
	"github.com/mecitsemerci/go-todo-app/internal/core/interfaces"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IdGenerator struct{}

func NewIdGenerator() interfaces.IDGenerator {
	return &IdGenerator{}
}

func (i *IdGenerator) NewID() domain.ID {
	return domain.ID(primitive.NewObjectID().Hex())
}

func (i *IdGenerator) IDFromString(str string) (domain.ID, error) {
	oid, err := primitive.ObjectIDFromHex(str)
	if err != nil {
		return domain.NilID, err
	}
	return domain.ID(oid.Hex()), err
}
