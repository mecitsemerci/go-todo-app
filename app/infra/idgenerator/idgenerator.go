package idgenerator

import (
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func ObjectIDFromID(id domain.ID) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id.String())
}

func NewID() domain.ID {
	return domain.ID(primitive.NewObjectIDFromTimestamp(time.Now().UTC()).Hex())
}

func IDFromObjectID(oid primitive.ObjectID) domain.ID {
	return domain.ID(oid.Hex())
}

func IDFromStr(str string) (domain.ID, error) {
	oid, err := primitive.ObjectIDFromHex(str)
	if err != nil {
		return domain.NilID, err
	}
	return domain.ID(oid.Hex()), nil
}
