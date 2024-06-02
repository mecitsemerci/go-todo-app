package service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"

	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
)

const (
	otelName = "Service"
)

// IDGenerator represents generic domain ID generator
type IDGenerator[T any] interface {
	NewID() T
	IDFromString(str string) (T, error)
}

// TodoRepository interface guides all operations
type TodoRepository interface {
	GetAll(ctx context.Context) ([]todo.Todo, error)
	GetByID(ctx context.Context, id todo.ID) (todo.Todo, error)
	Insert(ctx context.Context, todo todo.Todo) (todo.ID, error)
	Update(ctx context.Context, todo todo.Todo) error
	Delete(ctx context.Context, id todo.ID) error
	Close(ctx context.Context) error
}

// TodoService provide operations of TodoService interface
type TodoService struct {
	todoRepository TodoRepository
	idGenerator    IDGenerator[todo.ID]
}

// NewTodoService returns a new TodoService
func NewTodoService(todoRepo TodoRepository, idGen IDGenerator[todo.ID]) *TodoService {
	return &TodoService{
		todoRepository: todoRepo,
		idGenerator:    idGen,
	}
}

// GetAll returns all todo items
func (ts *TodoService) GetAll(ctx context.Context) ([]todo.Todo, error) {
	spanCtx, span := otel.Tracer(otelName).Start(ctx, "TodoService.GetAll")
	defer span.End()

	return ts.todoRepository.GetAll(spanCtx)
}

// Find todo item by given ID
func (ts *TodoService) Find(ctx context.Context, todoID todo.ID) (todo.Todo, error) {
	spanCtx, span := otel.Tracer(otelName).Start(ctx, "TodoService.Find")
	defer span.End()

	return ts.todoRepository.GetByID(spanCtx, todoID)
}

// Create todo item by given todo item
func (ts *TodoService) Create(ctx context.Context, t todo.Todo) (todo.ID, error) {
	spanCtx, span := otel.Tracer(otelName).Start(ctx, "TodoService.Create")
	defer span.End()

	// Set fields
	t.ID = ts.idGenerator.NewID()
	t.Completed = false
	t.CreatedAt = time.Now().UTC()
	t.UpdatedAt = time.Now().UTC()

	//Save
	id, err := ts.todoRepository.Insert(spanCtx, t)

	if err != nil {
		return todo.ZeroID, errors.Wrap(err, "insert failed")
	}
	return id, nil
}

// Update todo item by given todo item
func (ts *TodoService) Update(ctx context.Context, t todo.Todo) error {
	spanCtx, span := otel.Tracer(otelName).Start(ctx, "TodoService.Update")
	defer span.End()

	t.UpdatedAt = time.Now().UTC()

	return ts.todoRepository.Update(spanCtx, t)
}

// Delete todo item by given todo ID
func (ts *TodoService) Delete(ctx context.Context, tid todo.ID) error {
	spanCtx, span := otel.Tracer(otelName).Start(ctx, "TodoService.Delete")
	defer span.End()

	return ts.todoRepository.Delete(spanCtx, tid)
}
