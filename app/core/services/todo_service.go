package services

import (
	"github.com/google/uuid"
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/adapters/mementodb"
)

type IService interface {
	GetAll() ([]todo.Todo, error)
	Find(todoId int) (todo.Todo, error)
	Create(todo todo.Todo) (uuid.UUID, error)
	Update(todo todo.Todo) (bool, error)
	Delete(todoId int) (bool, error)
}
type TodoService struct {
	BaseService
	todoRepository mementodb.TodoAdapter
}

func (service *TodoService) GetAll() ([]todo.Todo, error) {
	items, err := service.todoRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return items, nil
}

func (service *TodoService) Find(todoId uuid.UUID) (todo.Todo, error) {
	entity, err := service.todoRepository.GetById(todoId)

	if err != nil {
		return todo.Todo{}, err
	}

	return entity, nil
}

func (service *TodoService) Create(todo todo.Todo) (uuid.UUID, error) {
	entityId, err := service.todoRepository.Insert(todo)

	if err != nil {
		return uuid.UUID{}, err
	}

	return entityId, nil
}

func (service *TodoService) Update(todo todo.Todo) (bool, error) {
	result, err := service.todoRepository.Update(todo)

	if err != nil {
		return false, err
	}

	return result, nil
}

func (service *TodoService) Delete(todoId uuid.UUID) (bool, error) {
	result, err := service.todoRepository.Delete(todoId)

	if err != nil {
		return false, err
	}

	return result, nil
}
