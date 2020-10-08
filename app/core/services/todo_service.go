package services

import (
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain"
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/datetime"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/idgenerator"
)

type ITodoService interface {
	GetAll() ([]*todo.Todo, error)
	Find(todoId domain.ID) (*todo.Todo, error)
	Create(todo todo.Todo) (domain.ID, error)
	Update(todo todo.Todo) error
	Delete(todoId domain.ID) error
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

func (service *TodoService) Find(todoId domain.ID) (*todo.Todo, error) {
	model, err := service.TodoRepository.GetById(todoId)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (service *TodoService) Create(todo todo.Todo) (domain.ID, error) {
	// Set fields
	todo.Id = idgenerator.NewID()
	todo.Completed = false
	todo.CreatedAt = datetime.Now()
	todo.UpdatedAt = datetime.Now()

	//Save
	id, err := service.TodoRepository.Insert(todo)
	if err != nil {
		return domain.NilID, err
	}
	return id, nil
}

func (service *TodoService) Update(todo todo.Todo) error {
	err := service.TodoRepository.Update(todo)
	return err
}

func (service *TodoService) Delete(todoId domain.ID) error {
	if err := service.TodoRepository.Delete(todoId); err != nil {
		return err
	}
	return nil
}
