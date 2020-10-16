package mongodb

import (
	"errors"
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain"
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/datetime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type TodoRepository struct {
	DbCtx DbContext
}

func NewTodoRepository(dbContext DbContext) *TodoRepository {
	return &TodoRepository{DbCtx: dbContext}
}

func (repo *TodoRepository) GetAll() ([]*todo.Todo, error) {
	var todos []*todo.Todo

	//Configure Find Query
	findOptions := options.Find()

	//Connect
	repo.DbCtx.Connect()

	cur, err := repo.DbCtx.TodoCollection.Find(repo.DbCtx.Context, bson.D{}, findOptions)

	if err != nil {
		return nil, err
	}

	for cur.Next(repo.DbCtx.Context) {
		var entity Todo
		err := cur.Decode(&entity)
		if err != nil {
			log.Println(err.Error())
		}
		todos = append(todos, entity.ToModel())
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	_ = cur.Close(repo.DbCtx.Context)

	//Disconnect
	defer repo.DbCtx.Disconnect()

	return todos, nil
}

func (repo *TodoRepository) GetById(id domain.ID) (*todo.Todo, error) {
	var t Todo
	oid, err := primitive.ObjectIDFromHex(id.String())
	if err != nil {
		return nil, err
	}
	//Filter
	filter := bson.M{"_id": bson.M{"$eq": oid}}

	//Connect
	repo.DbCtx.Connect()
	//Find Item by Id
	err = repo.DbCtx.TodoCollection.FindOne(repo.DbCtx.Context, filter).Decode(&t)

	//Disconnect
	defer repo.DbCtx.Disconnect()

	if err != nil {
		return nil, err
	}

	return t.ToModel(), nil
}

func (repo *TodoRepository) Insert(todo todo.Todo) (domain.ID, error) {
	// Data object
	var t Todo
	err := t.FromModel(&todo)
	if err != nil {
		return domain.NilID, err
	}
	//Connect
	repo.DbCtx.Connect()

	//Insert Item
	result, err := repo.DbCtx.TodoCollection.InsertOne(repo.DbCtx.Context, t)

	//Disconnect
	defer repo.DbCtx.Disconnect()

	if err != nil {
		return domain.NilID, err
	}

	// Return inserted item id
	var oid ObjectId
	oid.Set(result.InsertedID.(primitive.ObjectID).Hex())
	return &oid, nil
}

func (repo *TodoRepository) Update(todo todo.Todo) error {
	oid, err := primitive.ObjectIDFromHex(todo.Id.String())
	if err != nil {
		return err
	}
	//Filter
	filter := bson.M{"_id": bson.M{"$eq": oid}}

	//Update fields
	document := bson.M{"$set": bson.M{
		"title":       todo.Title,
		"description": todo.Description,
		"completed":   todo.Completed,
		"updated_at":  datetime.Now(),
	}}

	//Connect
	repo.DbCtx.Connect()

	//Update Item
	result, err := repo.DbCtx.TodoCollection.UpdateOne(repo.DbCtx.Context, filter, document)

	//Disconnect
	defer repo.DbCtx.Disconnect()

	if err != nil {
		return err
	}

	if result.MatchedCount > 0 {
		return nil
	}
	return errors.New("no items have been updated")
}

func (repo *TodoRepository) Delete(id domain.ID) error {
	oid, err := primitive.ObjectIDFromHex(id.String())
	if err != nil {
		return err
	}

	//Filter
	filter := bson.M{"_id": bson.M{"$eq": oid}}

	//Connect
	repo.DbCtx.Connect()

	//Delete Item
	result, err := repo.DbCtx.TodoCollection.DeleteOne(repo.DbCtx.Context, filter)

	//Disconnect
	defer repo.DbCtx.Disconnect()

	if err != nil {
		return err
	}

	if result.DeletedCount > 0 {
		return nil
	}
	return errors.New("no item has been deleted")
}
