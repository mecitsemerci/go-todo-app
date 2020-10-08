package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/constants"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TodoServiceTestSuite struct {
	suite.Suite
	MockTodoRepository *todo.MockTodoRepository
}

func TestTodoServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TodoServiceTestSuite))
}

func (s *TodoServiceTestSuite) SetupTest() {
	s.MockTodoRepository = &todo.MockTodoRepository{}
}

func (s *TodoServiceTestSuite) Test_NewJobService_Should_Return_Service() {
	// When
	service := NewTodoService(s.MockTodoRepository)

	// Then
	assert.NotNil(s.T(), service)
}

// region Get All
func (s *TodoServiceTestSuite) Test_Get_All_Should_Return_Todo_List_When_There_Are_Items() {
	//Given
	s.MockTodoRepository.MockGetAll = func() ([]*todo.Todo, error) {
		return []*todo.Todo{
			{
				Id:          utils.NewOID(),
				Title:       "Task 1",
				Description: "Task 1 Desc",
				Completed:   false,
				CreatedAt:   utils.UtcNow(),
				UpdatedAt:   utils.UtcNow(),
			},
			{
				Id:          utils.NewOID(),
				Title:       "Task 2",
				Description: "Task 2 Desc",
				Completed:   false,
				CreatedAt:   utils.UtcNow(),
				UpdatedAt:   utils.UtcNow(),
			},
		}, nil
	}

	// When
	service := NewTodoService(s.MockTodoRepository)
	todos, err := service.GetAll()

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), todos)
	assert.Equal(s.T(), 2, len(todos))
	assert.NotEqual(s.T(), uuid.Nil, todos[0].Id)
	assert.Equal(s.T(), "Task 1", todos[0].Title)
	assert.Equal(s.T(), "Task 1 Desc", todos[0].Description)
	assert.Equal(s.T(), false, todos[0].Completed)
	assert.True(s.T(), utils.UtcNow().After(todos[0].CreatedAt))
	assert.True(s.T(), utils.UtcNow().After(todos[0].UpdatedAt))
}

func (s *TodoServiceTestSuite) Test_Get_All_Delete_Should_Return_Nil_When_There_Are_No_Items() {
	//Given
	s.MockTodoRepository.MockGetAll = func() ([]*todo.Todo, error) {
		return nil, nil
	}

	// When
	service := NewTodoService(s.MockTodoRepository)
	todos, err := service.GetAll()

	// Then
	assert.Nil(s.T(), err)
	assert.Nil(s.T(), todos)
}

func (s *TodoServiceTestSuite) Test_Get_All_Delete_Should_Return_Error_When_Has_Connection_Error() {
	//Given
	errorMessage := errors.New("connection error")
	s.MockTodoRepository.MockGetAll = func() ([]*todo.Todo, error) {
		return nil, errorMessage
	}

	// When
	service := NewTodoService(s.MockTodoRepository)
	todos, err := service.GetAll()

	// Then
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errorMessage.Error(), err.Error())
	assert.Nil(s.T(), todos)
}

//endregion

// region Find
func (s *TodoServiceTestSuite) Test_Find_Should_Return_Todo_When_Given_Todo_Id() {
	//Given
	todoId := utils.NewOID()
	s.MockTodoRepository.MockGetById = func(id string) (*todo.Todo, error) {
		return &todo.Todo{
			Id:          todoId,
			Title:       "Task 1",
			Description: "Task 1 Desc",
			Completed:   false,
			CreatedAt:   utils.UtcNow(),
			UpdatedAt:   utils.UtcNow(),
		}, nil
	}

	// When
	service := NewTodoService(s.MockTodoRepository)
	id, err := service.Find(todoId.Hex())

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), id)

}

func (s *TodoServiceTestSuite) Test_Find_Should_Return_Error_When_No_Item_Given_Todo_Id() {
	//Given
	todoId := utils.NewOID()
	s.MockTodoRepository.MockGetById = func(id string) (*todo.Todo, error) {
		return nil, errors.New("item not found")
	}

	// When
	service := NewTodoService(s.MockTodoRepository)
	entity, err := service.Find(todoId.Hex())

	// Then
	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), entity)
}

//endregion

// region Create
func (s *TodoServiceTestSuite) Test_Create_Todo_Should_Return_Created_Todo_Id_When_Given_Todo_Entity() {
	//Given
	entity := todo.Todo{
		Title:       "Foo",
		Description: "Foo Description",
		Completed:   false,
	}
	oid := utils.NewOID()
	s.MockTodoRepository.MockInsert = func(todo todo.Todo) (string, error) {
		return oid.Hex(), nil
	}

	// When
	service := NewTodoService(s.MockTodoRepository)
	id, err := service.Create(entity)

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), id)
	assert.Equal(s.T(), oid.Hex(), id)
}

func (s *TodoServiceTestSuite) Test_Create_Todo_Should_Return_Error_When_Connection_Error_Occurs() {
	//Given
	entity := todo.Todo{
		Title:       "Foo",
		Description: "Foo Description",
		Completed:   false,
	}
	errorMessage := errors.New("database connection refused")
	s.MockTodoRepository.MockInsert = func(todo todo.Todo) (string, error) {
		return constants.EmptyString, errorMessage
	}

	// When
	service := NewTodoService(s.MockTodoRepository)
	id, err := service.Create(entity)

	// Then
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errorMessage.Error(), err.Error())
	assert.Equal(s.T(), constants.EmptyString, id)
}

//endregion

// region Update
func (s *TodoServiceTestSuite) Test_Update_Todo_Should_Return_No_Error_When_Given_Todo_Entity() {
	//Given
	entity := todo.Todo{
		Title:       "Foo",
		Description: "Foo Description",
		Completed:   false,
	}
	s.MockTodoRepository.MockUpdate = func(todo todo.Todo) error {
		return nil
	}

	// When
	service := NewTodoService(s.MockTodoRepository)
	err := service.Update(entity)

	// Then
	assert.Nil(s.T(), err)
}

func (s *TodoServiceTestSuite) Test_Update_Todo_Should_Return_Error_When_Item_Not_Found() {
	//Given
	entity := todo.Todo{
		Title:       "Foo",
		Description: "Foo Description",
		Completed:   false,
	}
	errorMessage := errors.New("item not found")
	s.MockTodoRepository.MockUpdate = func(todo todo.Todo) error {
		return errorMessage
	}

	// When
	service := NewTodoService(s.MockTodoRepository)
	err := service.Update(entity)

	// Then
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errorMessage.Error(), err.Error())
}

//endregion

// region Delete
func (s *TodoServiceTestSuite) Test_Delete_Should_Return_No_Error_When_Given_Todo_Id() {
	//Given
	oid := utils.NewOID()
	s.MockTodoRepository.MockDelete = func(id string) error {
		return nil
	}

	// When
	service := NewTodoService(s.MockTodoRepository)
	err := service.Delete(oid.Hex())

	// Then
	assert.Nil(s.T(), err)
}

func (s *TodoServiceTestSuite) Test_Delete_Should_Return_Error_When_Given_Item_Not_Found() {
	//Given
	oid := utils.NewOID()
	errorMessage := errors.New("item not found")
	s.MockTodoRepository.MockDelete = func(id string) error {
		return errorMessage
	}

	// When
	service := NewTodoService(s.MockTodoRepository)
	err := service.Delete(oid.Hex())

	// Then
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), errorMessage.Error(), err.Error())
}

//endregion
