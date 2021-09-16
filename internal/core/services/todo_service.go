package services

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
	"github.com/mecitsemerci/go-todo-app/internal/core/interfaces"
)

//TodoService provide operations of TodoService interface
type TodoService struct {
	todoRepository interfaces.TodoRepository
	idGenerator    interfaces.IDGenerator
}

//NewTodoService returns a new TodoService
func NewTodoService(todoRepository interfaces.TodoRepository,
	idGenerator interfaces.IDGenerator) *TodoService {
	return &TodoService{
		todoRepository: todoRepository,
		idGenerator:    idGenerator,
	}
}

//GetAll returns all todo items
func (srv *TodoService) GetAll(ctx context.Context) ([]todo.Todo, error) {
	span, spanContext := opentracing.StartSpanFromContext(ctx, "TodoService-GetAll")
	defer span.Finish()

	return srv.todoRepository.GetAll(spanContext)
}

//Find todo item by given ID
func (srv *TodoService) Find(ctx context.Context, todoID domain.ID) (todo.Todo, error) {
	span, spanContext := opentracing.StartSpanFromContext(ctx, "TodoService-Find")
	defer span.Finish()

	return srv.todoRepository.GetByID(spanContext, todoID)
}

//Create todo item by given todo item
func (srv *TodoService) Create(ctx context.Context, todo todo.Todo) (domain.ID, error) {
	span, spanContext := opentracing.StartSpanFromContext(ctx, "TodoService-Create")
	defer span.Finish()

	// Set fields
	todo.ID = srv.idGenerator.NewID()
	todo.Completed = false
	todo.CreatedAt = time.Now().UTC()
	todo.UpdatedAt = time.Now().UTC()

	//Save
	id, err := srv.todoRepository.Insert(spanContext, todo)

	if err != nil {
		return domain.ZeroID, errors.Wrap(err, "insert failed")
	}
	return id, nil
}

//Update todo item by given todo item
func (srv *TodoService) Update(ctx context.Context, todo todo.Todo) error {
	span, spanContext := opentracing.StartSpanFromContext(ctx, "TodoService-Update")
	defer span.Finish()

	todo.UpdatedAt = time.Now().UTC()

	return srv.todoRepository.Update(spanContext, todo)
}

//Delete todo item by given todo ID
func (srv *TodoService) Delete(ctx context.Context, todoID domain.ID) error {
	span, spanContext := opentracing.StartSpanFromContext(ctx, "TodoService-Delete")
	defer span.Finish()

	return srv.todoRepository.Delete(spanContext, todoID)
}
