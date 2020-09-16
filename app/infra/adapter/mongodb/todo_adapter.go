package mongodb

import (
	"errors"
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/utility"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type TodoAdapter struct {
	DBContext *DbContext
}

func (adapter *TodoAdapter) Init() *TodoAdapter {
	adapter.DBContext = new(DbContext)
	return adapter
}

func (adapter *TodoAdapter) GetById(id primitive.ObjectID) (*todo.Todo, error) {
	//Connect
	adapter.DBContext.Connect()
	var item todo.Todo
	filter := bson.M{"_id": bson.M{"$eq": id}}

	//Delete Item
	err := adapter.DBContext.TodoCollection.FindOne(adapter.DBContext.Context, filter).Decode(&item)

	//Disconnect
	defer adapter.DBContext.Disconnect()

	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (adapter *TodoAdapter) GetAll() ([]*todo.Todo, error) {
	var todos []*todo.Todo
	//Connect
	adapter.DBContext.Connect()
	findOptions := options.Find()

	//Insert Item
	cur, err := adapter.DBContext.TodoCollection.Find(adapter.DBContext.Context, bson.D{}, findOptions)

	if err != nil {
		return nil, err
	}

	for cur.Next(adapter.DBContext.Context) {
		var item todo.Todo
		err := cur.Decode(&item)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, &item)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	_ = cur.Close(adapter.DBContext.Context)

	//Disconnect
	defer adapter.DBContext.Disconnect()

	return todos, nil
}

func (adapter *TodoAdapter) Insert(entity todo.Todo) (primitive.ObjectID, error) {

	entity.Id = primitive.NewObjectIDFromTimestamp(utility.UtcNow())
	entity.CreatedAt = utility.UtcNow()
	entity.UpdatedAt = utility.UtcNow()

	//Connect
	adapter.DBContext.Connect()

	//Insert Item
	result, err := adapter.DBContext.TodoCollection.InsertOne(adapter.DBContext.Context, entity)

	//Disconnect
	defer adapter.DBContext.Disconnect()

	if err != nil {
		return primitive.NilObjectID, err
	}

	// Return inserted item id
	return result.InsertedID.(primitive.ObjectID), nil
}

func (adapter *TodoAdapter) Update(entity todo.Todo) error {
	//Connect
	adapter.DBContext.Connect()

	filter := bson.M{"_id": bson.M{"$eq": entity.Id}}
	update := bson.M{"$set": bson.M{
		"title":       entity.Title,
		"description": entity.Description,
		"is_done":     entity.IsDone,
		"updated_at":  utility.UtcNow(),
	}}

	//Update Item
	result, err := adapter.DBContext.TodoCollection.UpdateOne(adapter.DBContext.Context, filter, update)

	//Disconnect
	defer adapter.DBContext.Disconnect()

	if err != nil {
		return err
	}

	if result.MatchedCount > 0 {
		return nil
	}
	return errors.New("no items have been updated")
}

func (adapter *TodoAdapter) Delete(id primitive.ObjectID) error {
	//Connect
	adapter.DBContext.Connect()

	filter := bson.M{"_id": bson.M{"$eq": id}}

	//Delete Item
	result, err := adapter.DBContext.TodoCollection.DeleteOne(adapter.DBContext.Context, filter)

	//Disconnect
	defer adapter.DBContext.Disconnect()

	if err != nil {
		return err
	}

	if result.DeletedCount > 0 {
		return nil
	}
	return errors.New("no item has been deleted")
}
