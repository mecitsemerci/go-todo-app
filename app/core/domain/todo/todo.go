package todo

import (
	"github.com/google/uuid"
	"time"
)

type Todo struct {
	Id        uuid.UUID
	Title     string
	IsDone    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
