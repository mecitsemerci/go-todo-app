package todo

import (
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

func (t *Todo) ID() string {
	return t.Id.Hex()
}
