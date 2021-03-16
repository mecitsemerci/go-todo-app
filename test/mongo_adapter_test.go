package test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/mecitsemerci/go-todo-app/internal/config"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
	"github.com/mecitsemerci/go-todo-app/internal/core/interfaces"
	"github.com/mecitsemerci/go-todo-app/internal/pkg/mongodb"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbname  = "TestTodoDb"
	timeout = 5
)

var (
	mongoClient *mongo.Client
	todoAdapter interfaces.TodoRepository
	idGenerator interfaces.IDGenerator
)

//region TestSetup
func setup() {
	config.Load()

	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	if client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoConfig.MongoURL)); err != nil {
		log.Fatalf("provide mongodb: %s", err.Error())
	} else {
		mongoClient = client
	}
	todoAdapter = mongodb.NewTodoAdapter(mongoClient, dbname)
	idGenerator = mongodb.NewIDGenerator()
}

func shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	_ = mongoClient.Database(dbname).Drop(ctx)
	if err := mongoClient.Disconnect(ctx); err != nil {
		log.Println(err.Error())
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func clear() {
	_ = mongoClient.Database(dbname).Collection(mongodb.CollectionName).Drop(context.Background())
}

func createFakeTodo() todo.Todo {
	return todo.Todo{
		ID:          idGenerator.NewID(),
		Title:       gofakeit.JobTitle(),
		Description: gofakeit.JobDescriptor(),
		Completed:   true,
		Priority:    1,
		CreatedAt:   gofakeit.Date().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}

//endregion

func Test_GetAll_Should_Return_Empty_Array_When_There_Is_No_Items(t *testing.T) {
	//Given
	clear()

	//When
	items, err := todoAdapter.GetAll(context.TODO())

	//Then
	assert.Nil(t, err)
	assert.Empty(t, items)
}

func Test_GetAll_Should_Return_A_Item_When_There_Is_A_Item(t *testing.T) {
	//Given
	clear()
	item := createFakeTodo()
	id, err := todoAdapter.Insert(context.TODO(), item)
	assert.Nil(t, err)
	assert.NotEmpty(t, id)

	//When
	items, err := todoAdapter.GetAll(context.TODO())

	//Then
	assert.Nil(t, err)
	assert.NotEmpty(t, items)
	assert.Equal(t, 1, len(items))
	assert.Equal(t, item.ID, items[0].ID)
	assert.Equal(t, item.Title, items[0].Title)
	assert.Equal(t, item.Description, items[0].Description)
	assert.Equal(t, item.Completed, items[0].Completed)
	assert.Equal(t, item.Priority, items[0].Priority)
	assert.Equal(t, item.CreatedAt.Format(time.RFC3339), items[0].CreatedAt.Format(time.RFC3339))
	assert.Equal(t, item.UpdatedAt.Format(time.RFC3339), items[0].UpdatedAt.Format(time.RFC3339))

}

func Test_GetAll_Should_Return_All_Items(t *testing.T) {
	//Given
	clear()
	count := 10
	for i := 0; i < count; i++ {
		item := createFakeTodo()
		id, err := todoAdapter.Insert(context.TODO(), item)
		assert.Nil(t, err)
		assert.NotEmpty(t, id)
	}

	//When
	items, err := todoAdapter.GetAll(context.TODO())

	//Then
	assert.Nil(t, err)
	assert.NotEmpty(t, items)
	assert.Equal(t, count, len(items))
}

func Test_GetById_Should_Return_Item_When_Given_Item_ID(t *testing.T) {
	//Given
	clear()
	fakeItem := createFakeTodo()
	id, err := todoAdapter.Insert(context.TODO(), fakeItem)
	assert.Nil(t, err)
	assert.NotEmpty(t, id)

	//When
	item, err := todoAdapter.GetByID(context.TODO(), id)

	//Then
	assert.Nil(t, err)
	assert.NotEmpty(t, item)
	assert.Equal(t, fakeItem.ID, item.ID)
	assert.Equal(t, fakeItem.Title, item.Title)
	assert.Equal(t, fakeItem.Description, item.Description)
	assert.Equal(t, fakeItem.Completed, item.Completed)
	assert.Equal(t, fakeItem.Priority, item.Priority)
	assert.Equal(t, fakeItem.CreatedAt.Format(time.RFC3339), item.CreatedAt.Format(time.RFC3339))
	assert.Equal(t, fakeItem.UpdatedAt.Format(time.RFC3339), item.UpdatedAt.Format(time.RFC3339))
}

func Test_GetById_Should_Return_Error_When_Given_Item_ID_Empty(t *testing.T) {
	//Given

	//When
	item, err := todoAdapter.GetByID(context.TODO(), domain.ZeroID)

	//Then
	assert.NotNil(t, err)
	assert.Empty(t, item)
}

func Test_GetById_Should_Return_Error_When_Given_Item_ID_Not_Exist(t *testing.T) {
	//Given

	//When
	item, err := todoAdapter.GetByID(context.TODO(), idGenerator.NewID())

	//Then
	assert.NotNil(t, err)
	assert.Empty(t, item)
}

func Test_Insert_Should_Return_ObjectId_When_Item_Created(t *testing.T) {
	//Given
	todoItem := createFakeTodo()

	//When
	id, err := todoAdapter.Insert(context.TODO(), todoItem)

	//Then
	assert.Nil(t, err)
	assert.Equal(t, todoItem.ID, id)
}

func Test_Insert_Should_Return_Error_When_No_Given_Item_ID(t *testing.T) {
	//Given
	todoItem := createFakeTodo()
	todoItem.ID = domain.ZeroID

	//When
	id, err := todoAdapter.Insert(context.TODO(), todoItem)

	//Then
	assert.NotNil(t, err)
	assert.Empty(t, id)
}

func Test_Insert_Should_Return_Error_When_Db_Connection_Not_Exist(t *testing.T) {
	//Given
	todoItem := createFakeTodo()
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoConfig.MongoURL))
	assert.Nil(t, err)
	adapter := mongodb.NewTodoAdapter(client, dbname)
	err = client.Disconnect(ctx)
	assert.Nil(t, err)

	//When
	id, err := adapter.Insert(context.TODO(), todoItem)

	//Then
	assert.NotNil(t, err)
	assert.Empty(t, id)
}

func Test_Update_Should_Return_No_Error_When_Item_Updated(t *testing.T) {
	//Given
	clear()
	todoItem := createFakeTodo()
	id, err := todoAdapter.Insert(context.TODO(), todoItem)
	assert.Nil(t, err)
	assert.NotEmpty(t, id)

	//When
	todoItem.Description = "Updated"
	err = todoAdapter.Update(context.TODO(), todoItem)

	//Then
	assert.Nil(t, err)
}

func Test_Update_Should_Return_Error_When_Item_Not_Exist(t *testing.T) {
	//Given
	todoItem := createFakeTodo()

	//When
	todoItem.Description = "Updated"
	err := todoAdapter.Update(context.TODO(), todoItem)

	//Then
	assert.NotNil(t, err)
}
