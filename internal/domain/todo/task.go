package todo

import (
	"time"

	"github.com/mecitsemerci/go-todo-app/internal/domain/auth"
	"github.com/mecitsemerci/go-todo-app/internal/domain/priority"
)

type Task struct {
	ID          TaskID
	Title       string
	Description string
	Completed   bool
	Priority    priority.Level
	Deleted     bool
	CreatedBy   auth.UserID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTask(title, description string, priority priority.Level, userID auth.UserID) *Task {
	now := time.Now().UTC()
	return &Task{
		ID:          NewTaskID(),
		Title:       title,
		Description: description,
		Priority:    priority,
		Completed:   false,
		CreatedBy:   userID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (t *Task) UpdateTitle(title string) *Task {
	t.Title = title
	t.UpdatedAt = time.Now().UTC()
	return t
}

func (t *Task) UpdateDescription(description string) *Task {
	t.Description = description
	t.UpdatedAt = time.Now().UTC()
	return t
}

func (t *Task) UpdatePriority(priority priority.Level) *Task {
	t.Priority = priority
	t.UpdatedAt = time.Now().UTC()
	return t
}

func (t *Task) UpdateCompleted(completed bool) *Task {
	t.Completed = completed
	t.UpdatedAt = time.Now().UTC()
	return t
}

func (t *Task) Delete() *Task {
	t.Deleted = true
	t.UpdatedAt = time.Now().UTC()
	return t
}

func (t *Task) Restore() *Task {
	t.Deleted = false
	t.UpdatedAt = time.Now().UTC()
	return t
}
