package services

import (
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/adapter/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ITodoService interface {
	GetAll() ([]*todo.Todo, error)
	Find(todoId string) (*todo.Todo, error)
	Create(todo todo.Todo) (string, error)
	Update(todo todo.Todo) error
	Delete(todoId string) error
}
type TodoService struct {
	BaseService
	todoRepository *mongodb.TodoAdapter
}

func (service *TodoService) Init() *TodoService {
	service.todoRepository = new(mongodb.TodoAdapter).Init()
	return service
}

func (service *TodoService) GetAll() ([]*todo.Todo, error) {
	items, err := service.todoRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (service *TodoService) Find(todoId string) (*todo.Todo, error) {
	oid, err := primitive.ObjectIDFromHex(todoId)
	if err != nil {
		return nil, err
	}
	entity, err := service.todoRepository.GetById(oid)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (service *TodoService) Create(todo todo.Todo) (string, error) {
	oid, err := service.todoRepository.Insert(todo)
	if err != nil {
		return "", err
	}
	return oid.Hex(), nil
}

func (service *TodoService) Update(todo todo.Todo) error {
	err := service.todoRepository.Update(todo)
	return err
}

func (service *TodoService) Delete(todoId string) error {

	oid, err := primitive.ObjectIDFromHex(todoId)
	if err != nil {
		return err
	}

	err = service.todoRepository.Delete(oid)
	if err != nil {
		return err
	}
	return nil
}
