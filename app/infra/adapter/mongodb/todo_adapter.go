package mongodb

import (
	"errors"
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/constants"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type TodoAdapter struct {
	DbCtx DbContext
}

func NewTodoAdapter(dbContext DbContext) *TodoAdapter {
	return &TodoAdapter{DbCtx: dbContext}
}

func (adapter *TodoAdapter) GetAll() ([]*todo.Todo, error) {
	var todos []*todo.Todo

	//Configure Find Query
	findOptions := options.Find()

	//Connect
	adapter.DbCtx.Connect()

	cur, err := adapter.DbCtx.TodoCollection.Find(adapter.DbCtx.Context, bson.D{}, findOptions)

	if err != nil {
		return nil, err
	}

	for cur.Next(adapter.DbCtx.Context) {
		var item todo.Todo
		err := cur.Decode(&item)
		if err != nil {
			log.Println(err.Error())
		}
		todos = append(todos, &item)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	_ = cur.Close(adapter.DbCtx.Context)

	//Disconnect
	defer adapter.DbCtx.Disconnect()

	return todos, nil
}

func (adapter *TodoAdapter) GetById(id string) (*todo.Todo, error) {
	var item todo.Todo

	oid, err := utils.OIDFromStr(id)

	if err != nil {
		return nil, err
	}

	//Filter
	filter := bson.M{"_id": bson.M{"$eq": oid}}

	//Connect
	adapter.DbCtx.Connect()

	//Find Item by Id
	err = adapter.DbCtx.TodoCollection.FindOne(adapter.DbCtx.Context, filter).Decode(&item)

	//Disconnect
	defer adapter.DbCtx.Disconnect()

	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (adapter *TodoAdapter) Insert(todo todo.Todo) (string, error) {
	// Set Fields
	todo.Id = utils.NewOID()
	todo.CreatedAt = utils.UtcNow()
	todo.UpdatedAt = utils.UtcNow()
	todo.Completed = false
	//Connect
	adapter.DbCtx.Connect()

	//Insert Item
	result, err := adapter.DbCtx.TodoCollection.InsertOne(adapter.DbCtx.Context, todo)

	//Disconnect
	defer adapter.DbCtx.Disconnect()

	if err != nil {
		return constants.EmptyString, err
	}

	// Return inserted item id
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (adapter *TodoAdapter) Update(todo todo.Todo) error {
	//Filter
	filter := bson.M{"_id": bson.M{"$eq": todo.Id}}

	//Update fields
	update := bson.M{"$set": bson.M{
		"title":       todo.Title,
		"description": todo.Description,
		"completed":   todo.Completed,
		"updated_at":  utils.UtcNow(),
	}}

	//Connect
	adapter.DbCtx.Connect()

	//Update Item
	result, err := adapter.DbCtx.TodoCollection.UpdateOne(adapter.DbCtx.Context, filter, update)

	//Disconnect
	defer adapter.DbCtx.Disconnect()

	if err != nil {
		return err
	}

	if result.MatchedCount > 0 {
		return nil
	}
	return errors.New("no items have been updated")
}

func (adapter *TodoAdapter) Delete(id string) error {
	oid, err := utils.OIDFromStr(id)

	if err != nil {
		return err
	}

	//Filter
	filter := bson.M{"_id": bson.M{"$eq": oid}}

	//Connect
	adapter.DbCtx.Connect()

	//Delete Item
	result, err := adapter.DbCtx.TodoCollection.DeleteOne(adapter.DbCtx.Context, filter)

	//Disconnect
	defer adapter.DbCtx.Disconnect()

	if err != nil {
		return err
	}

	if result.DeletedCount > 0 {
		return nil
	}
	return errors.New("no item has been deleted")
}
