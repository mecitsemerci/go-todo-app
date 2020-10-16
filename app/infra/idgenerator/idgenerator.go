package idgenerator

import (
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/repository/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func NewID() domain.ID {
	var oid mongodb.ObjectId
	oid.Set(primitive.NewObjectIDFromTimestamp(time.Now().UTC()).Hex())
	return &oid
}

func IDFromString(str string) domain.ID {
	var oid mongodb.ObjectId
	oid.Set(str)
	return &oid
}
