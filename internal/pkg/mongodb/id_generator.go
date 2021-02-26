package mongodb

import (
	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//IDGenerator for mongodb
type IDGenerator struct{}

//NewIDGenerator returns mongodb IDGenerator
func NewIDGenerator() *IDGenerator {
	return &IDGenerator{}
}

//NewID returns domain ID for mongo objectID
func (i *IDGenerator) NewID() domain.ID {
	return domain.ID(primitive.NewObjectID().Hex())
}

//IDFromString converts string id to domain ID for mongo objectID
func (i *IDGenerator) IDFromString(str string) (domain.ID, error) {
	oid, err := primitive.ObjectIDFromHex(str)
	if err != nil {
		return domain.ZeroID, err
	}
	return domain.ID(oid.Hex()), err
}
