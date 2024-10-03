package todo

import "github.com/google/uuid"

type TaskID string

func (t TaskID) String() string {
	return string(t)
}

func (t TaskID) Equals(other TaskID) bool {
	return t.String() == other.String()
}

func NewTaskID() TaskID {
	return TaskID(uuid.New().String())
}
