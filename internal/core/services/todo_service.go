package services

import (
	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
	"github.com/mecitsemerci/go-todo-app/internal/core/interfaces"
	"time"
)

type TodoService struct {
	todoRepository interfaces.TodoRepository
	idGenerator    interfaces.IDGenerator
}

func NewTodoService(todoRepository interfaces.TodoRepository,
	idGenerator interfaces.IDGenerator) interfaces.TodoService {
	return &TodoService{
		todoRepository: todoRepository,
		idGenerator:    idGenerator,
	}
}

func (srv *TodoService) GetAll() ([]todo.Todo, error) {
	return srv.todoRepository.GetAll()
}

func (srv *TodoService) Find(todoId domain.ID) (*todo.Todo, error) {
	return srv.todoRepository.GetById(todoId)
}

func (srv *TodoService) Create(todo todo.Todo) (domain.ID, error) {
	// Set fields
	todo.ID = srv.idGenerator.NewID()
	todo.Completed = false
	todo.CreatedAt = time.Now().UTC()
	todo.UpdatedAt = time.Now().UTC()

	//Save
	id, err := srv.todoRepository.Insert(todo)

	if err != nil {
		return domain.NilID, err
	}
	return id, nil
}

func (srv *TodoService) Update(todo todo.Todo) error {
	todo.UpdatedAt = time.Now().UTC()

	return srv.todoRepository.Update(todo)
}

func (srv *TodoService) Delete(todoId domain.ID) error {
	return srv.todoRepository.Delete(todoId)
}
