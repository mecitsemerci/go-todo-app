package todo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Todo struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	IsDone      bool               `bson:"is_done,omitempty"`
	CreatedAt   time.Time          `bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `bson:"updated_at,omitempty"`
}
