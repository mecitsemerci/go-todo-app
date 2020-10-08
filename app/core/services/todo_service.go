package services

import (
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/constants"
)

type ITodoService interface {
	GetAll() ([]*todo.Todo, error)
	Find(todoId string) (*todo.Todo, error)
	Create(todo todo.Todo) (string, error)
	Update(todo todo.Todo) error
	Delete(todoId string) error
}
type TodoService struct {
	TodoRepository todo.ITodoRepository
}

func NewTodoService(todoRepository todo.ITodoRepository) *TodoService {
	return &TodoService{TodoRepository: todoRepository}
}

func (service *TodoService) GetAll() ([]*todo.Todo, error) {
	items, err := service.TodoRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (service *TodoService) Find(todoId string) (*todo.Todo, error) {
	entity, err := service.TodoRepository.GetById(todoId)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (service *TodoService) Create(todo todo.Todo) (string, error) {
	id, err := service.TodoRepository.Insert(todo)
	if err != nil {
		return constants.EmptyString, err
	}
	return id, nil
}

func (service *TodoService) Update(todo todo.Todo) error {
	err := service.TodoRepository.Update(todo)
	return err
}

func (service *TodoService) Delete(todoId string) error {
	if err := service.TodoRepository.Delete(todoId); err != nil {
		return err
	}
	return nil
}
