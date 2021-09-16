package services

import (
	"context"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
	"github.com/mecitsemerci/go-todo-app/internal/core/mocks"
)

type TodoServiceTestSuite struct {
	suite.Suite
	MockTodoRepository *mocks.MockTodoRepository
	MockIDGenerator    *mocks.MockIDGenerator
}

func TestTodoServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TodoServiceTestSuite))
}

func (s *TodoServiceTestSuite) SetupTest() {
	s.MockTodoRepository = &mocks.MockTodoRepository{}
}

func (s *TodoServiceTestSuite) Test_NewJobService_Should_Return_Service() {
	// When
	service := NewTodoService(s.MockTodoRepository, s.MockIDGenerator)

	// Then
	assert.NotNil(s.T(), service)
}

// region Get All
func (s *TodoServiceTestSuite) Test_Get_All_Should_Return_Todo_List_When_There_Are_Items() {
	//Given
	s.MockTodoRepository.MockGetAll = func(ctx context.Context) ([]todo.Todo, error) {
		return []todo.Todo{
			{
				ID:          s.MockIDGenerator.NewID(),
				Title:       "Task 1",
				Description: "Task 1 Desc",
				Completed:   false,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				ID:          s.MockIDGenerator.NewID(),
				Title:       "Task 2",
				Description: "Task 2 Desc",
				Completed:   false,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		}, nil
	}

	// When
	service := NewTodoService(s.MockTodoRepository, s.MockIDGenerator)
	todos, err := service.GetAll(context.TODO())

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), todos)
	assert.Equal(s.T(), 2, len(todos))
	assert.NotEqual(s.T(), domain.ZeroID, todos[0].ID)
	assert.Equal(s.T(), "Task 1", todos[0].Title)
	assert.Equal(s.T(), "Task 1 Desc", todos[0].Description)
	assert.Equal(s.T(), false, todos[0].Completed)
	assert.True(s.T(), time.Now().After(todos[0].CreatedAt))
	assert.True(s.T(), time.Now().After(todos[0].UpdatedAt))
}

func (s *TodoServiceTestSuite) Test_Get_All_Delete_Should_Return_Nil_When_There_Are_No_Items() {
	//Given
	s.MockTodoRepository.MockGetAll = func(ctx context.Context) ([]todo.Todo, error) {
		return nil, nil
	}

	// When
	service := NewTodoService(s.MockTodoRepository, s.MockIDGenerator)
	todos, err := service.GetAll(context.TODO())

	// Then
	assert.Nil(s.T(), err)
	assert.Nil(s.T(), todos)
}

func (s *TodoServiceTestSuite) Test_Get_All_Delete_Should_Return_Error_When_Has_Connection_Error() {
	//Given
	errorMessage := errors.New("connection error")
	s.MockTodoRepository.MockGetAll = func(ctx context.Context) ([]todo.Todo, error) {
		return nil, errorMessage
	}

	// When
	service := NewTodoService(s.MockTodoRepository, s.MockIDGenerator)
	todos, err := service.GetAll(context.TODO())

	// Then
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errorMessage.Error(), err.Error())
	assert.Nil(s.T(), todos)
}

//endregion

// region Find
func (s *TodoServiceTestSuite) Test_Find_Should_Return_Todo_When_Given_Todo_Id() {
	//Given
	todoID := s.MockIDGenerator.NewID()
	s.MockTodoRepository.MockGetByID = func(ctx context.Context, id domain.ID) (todo.Todo, error) {
		return todo.Todo{
			ID:          todoID,
			Title:       "Task 1",
			Description: "Task 1 Desc",
			Completed:   false,
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
		}, nil
	}

	// When
	service := NewTodoService(s.MockTodoRepository, s.MockIDGenerator)
	t, err := service.Find(context.TODO(), todoID)

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), t)
	assert.NotEqual(s.T(), t.ID, domain.ZeroID)
	assert.Equal(s.T(), todoID, t.ID)
}

func (s *TodoServiceTestSuite) Test_Find_Should_Return_Error_When_No_Item_Given_Todo_Id() {
	//Given
	todoID := s.MockIDGenerator.NewID()
	s.MockTodoRepository.MockGetByID = func(ctx context.Context, id domain.ID) (todo.Todo, error) {
		return todo.DefaultTodo, errors.New("item not found")
	}

	// When
	service := NewTodoService(s.MockTodoRepository, s.MockIDGenerator)
	t, err := service.Find(context.TODO(), todoID)

	// Then
	assert.NotNil(s.T(), err)
	assert.Empty(s.T(), t)
}

//endregion

// region Create
func (s *TodoServiceTestSuite) Test_Create_Todo_Should_Return_Created_Todo_Id_When_Given_Todo_Entity() {
	//Given
	model := todo.Todo{
		Title:       "Foo",
		Description: "Foo Description",
		Completed:   false,
	}
	newID := s.MockIDGenerator.NewID()
	s.MockTodoRepository.MockInsert = func(ctx context.Context, todo todo.Todo) (domain.ID, error) {
		return newID, nil
	}

	// When
	service := NewTodoService(s.MockTodoRepository, s.MockIDGenerator)
	modelID, err := service.Create(context.TODO(), model)

	// Then
	assert.Nil(s.T(), err)
	assert.NotEqual(s.T(), modelID, domain.ZeroID)
	assert.Equal(s.T(), newID, modelID)
}

func (s *TodoServiceTestSuite) Test_Create_Todo_Should_Return_Error_When_Connection_Error_Occurs() {
	//Given
	model := todo.Todo{
		Title:       "Foo",
		Description: "Foo Description",
		Completed:   false,
	}
	errorMessage := errors.New("database connection refused")
	s.MockTodoRepository.MockInsert = func(ctx context.Context, todo todo.Todo) (domain.ID, error) {
		return domain.ZeroID, errorMessage
	}

	// When
	service := NewTodoService(s.MockTodoRepository, s.MockIDGenerator)
	id, err := service.Create(context.TODO(), model)

	// Then
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errors.Wrap(errorMessage, "insert failed").Error(), err.Error())
	assert.Equal(s.T(), domain.ZeroID, id)
}

//endregion

// region Update
func (s *TodoServiceTestSuite) Test_Update_Todo_Should_Return_No_Error_When_Given_Todo_Entity() {
	//Given
	model := todo.Todo{
		Title:       "Foo",
		Description: "Foo Description",
		Completed:   false,
	}
	s.MockTodoRepository.MockUpdate = func(ctx context.Context, todo todo.Todo) error {
		return nil
	}

	// When
	service := NewTodoService(s.MockTodoRepository, s.MockIDGenerator)
	err := service.Update(context.TODO(), model)

	// Then
	assert.Nil(s.T(), err)
}

func (s *TodoServiceTestSuite) Test_Update_Todo_Should_Return_Error_When_Item_Not_Found() {
	//Given
	model := todo.Todo{
		Title:       "Foo",
		Description: "Foo Description",
		Completed:   false,
	}
	errorMessage := errors.New("item not found")
	s.MockTodoRepository.MockUpdate = func(ctx context.Context, todo todo.Todo) error {
		return errorMessage
	}

	// When
	service := NewTodoService(s.MockTodoRepository, s.MockIDGenerator)
	err := service.Update(context.TODO(), model)

	// Then
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errorMessage.Error(), err.Error())
}

//endregion

// region Delete
func (s *TodoServiceTestSuite) Test_Delete_Should_Return_No_Error_When_Given_Todo_Id() {
	//Given
	newID := s.MockIDGenerator.NewID()
	s.MockTodoRepository.MockDelete = func(ctx context.Context, id domain.ID) error {
		return nil
	}

	// When
	service := NewTodoService(s.MockTodoRepository, s.MockIDGenerator)
	err := service.Delete(context.TODO(), newID)

	// Then
	assert.Nil(s.T(), err)
}

func (s *TodoServiceTestSuite) Test_Delete_Should_Return_Error_When_Given_Item_Not_Found() {
	//Given
	newID := s.MockIDGenerator.NewID()
	errorMessage := errors.New("item not found")
	s.MockTodoRepository.MockDelete = func(ctx context.Context, id domain.ID) error {
		return errorMessage
	}

	// When
	service := NewTodoService(s.MockTodoRepository, s.MockIDGenerator)
	err := service.Delete(context.TODO(), newID)

	// Then
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errorMessage.Error(), err.Error())
}

//endregion
