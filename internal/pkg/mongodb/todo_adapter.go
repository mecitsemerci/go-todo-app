package mongodb

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
	"github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	//CollectionName is MongoDB items collection name
	CollectionName = "todos"
)

var (
	//ErrNoItemsUpdated has no items updated errormessage
	ErrNoItemsUpdated = errors.New("mongodb: no items have been updated")

	//ErrNoItemsDeleted has no items deleted error message
	ErrNoItemsDeleted = errors.New("mongodb: no items have been deleted")
)

//TodoAdapter is a mongodb implementation of TodoRepository
type TodoAdapter struct {
	client     *mongo.Client
	collection *mongo.Collection
}

//NewTodoAdapter returns a mongodb adapter according to TodoRepository
func NewTodoAdapter(client *mongo.Client, dbName string) *TodoAdapter {
	return &TodoAdapter{
		client:     client,
		collection: client.Database(dbName).Collection(CollectionName),
	}
}

// GetAll returns all items
func (a *TodoAdapter) GetAll(ctx context.Context) ([]todo.Todo, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "TodoAdapter-GetAll")
	defer span.Finish()

	findOptions := options.Find()

	cur, err := a.collection.Find(ctx, bson.D{}, findOptions)

	var todos = make([]todo.Todo, 0)

	if err != nil {
		return todos, err
	}

	var errs = make([]string, 0)

	for cur.Next(ctx) {
		var entity Todo
		err := cur.Decode(&entity)
		if err != nil {
			errs = append(errs, fmt.Sprintf("decode error:%s", err.Error()))
		}
		todos = append(todos, entity.ToModel())
	}

	if err := cur.Err(); err != nil {
		errs = append(errs, fmt.Sprintf("cursor error:%s", err.Error()))
	}

	_ = cur.Close(ctx)

	if len(errs) > 0 {
		return todos, errors.New(strings.Join(errs, ";"))
	}

	return todos, nil
}

//GetByID retrieve item according to given item ID
func (a *TodoAdapter) GetByID(ctx context.Context, id domain.ID) (todo.Todo, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "TodoAdapter-GetByID")
	defer span.Finish()

	var t Todo

	oid, err := primitive.ObjectIDFromHex(string(id))

	if err != nil {
		return todo.DefaultTodo, err
	}

	filter := bson.M{"_id": bson.M{"$eq": oid}}

	err = a.collection.FindOne(ctx, filter).Decode(&t)

	if err != nil {
		return todo.DefaultTodo, err
	}

	return t.ToModel(), nil
}

//Insert item according to given item
func (a *TodoAdapter) Insert(ctx context.Context, t todo.Todo) (domain.ID, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "TodoAdapter-Insert")
	defer span.Finish()

	var todoItem Todo

	err := todoItem.FromModel(&t)

	if err != nil {
		return domain.ZeroID, err
	}

	result, err := a.collection.InsertOne(ctx, bson.D{
		primitive.E{Key: "_id", Value: todoItem.ID},
		primitive.E{Key: "title", Value: todoItem.Title},
		primitive.E{Key: "description", Value: todoItem.Description},
		primitive.E{Key: "priority_level", Value: todoItem.Priority},
		primitive.E{Key: "completed", Value: todoItem.Completed},
		primitive.E{Key: "created_at", Value: todoItem.CreatedAt},
		primitive.E{Key: "updated_at", Value: todoItem.UpdatedAt},
	})

	if err != nil {
		return domain.ZeroID, err
	}

	return domain.ID(result.InsertedID.(primitive.ObjectID).Hex()), nil
}

//Update item according to given item
func (a *TodoAdapter) Update(ctx context.Context, t todo.Todo) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "TodoAdapter-Update")
	defer span.Finish()

	oid, err := primitive.ObjectIDFromHex(string(t.ID))

	if err != nil {
		return err
	}

	//Filter
	filter := bson.M{"_id": bson.M{"$eq": oid}}

	//Update fields
	document := bson.M{"$set": bson.M{
		"title":          t.Title,
		"description":    t.Description,
		"priority_level": t.Priority,
		"completed":      t.Completed,
		"updated_at":     t.UpdatedAt,
	}}

	//Update Item
	result, err := a.collection.UpdateOne(ctx, filter, document)

	if err != nil {
		return err
	}

	if result.MatchedCount > 0 {
		return nil
	}

	return ErrNoItemsUpdated
}

//Delete item according to given todoID
func (a *TodoAdapter) Delete(ctx context.Context, id domain.ID) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "TodoAdapter-Delete")
	defer span.Finish()

	oid, err := primitive.ObjectIDFromHex(string(id))

	if err != nil {
		return err
	}

	filter := bson.M{"_id": bson.M{"$eq": oid}}

	result, err := a.collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if result.DeletedCount > 0 {
		return nil
	}
	return ErrNoItemsDeleted
}

//Close disconnects the connection
func (a *TodoAdapter) Close(ctx context.Context) error {
	return a.client.Disconnect(ctx)
}
