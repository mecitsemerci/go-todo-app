package mongodb

import (
	"context"
	"errors"
	"fmt"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

const (
	CollectionName = "todos"
	Timeout        = 5
)

var (
	ErrNoItemsUpdated = errors.New("mongodb: no items have been updated")
	ErrNoItemsDeleted = errors.New("mongodb: no items have been deleted")
)

type TodoAdapter struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewTodoAdapter(client *mongo.Client, dbName string) *TodoAdapter {
	return &TodoAdapter{
		client:     client,
		collection: client.Database(dbName).Collection(CollectionName),
	}
}

func (repo *TodoAdapter) GetAll() ([]todo.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	findOptions := options.Find()

	cur, err := repo.collection.Find(ctx, bson.D{}, findOptions)

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
		todos = append(todos, *entity.ToModel())
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

func (repo *TodoAdapter) GetById(id domain.ID) (*todo.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	var t Todo

	oid, err := primitive.ObjectIDFromHex(string(id))

	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": bson.M{"$eq": oid}}

	err = repo.collection.FindOne(ctx, filter).Decode(&t)

	if err != nil {
		return nil, err
	}

	return t.ToModel(), nil
}

func (repo *TodoAdapter) Insert(t todo.Todo) (domain.ID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	var todoItem Todo

	err := todoItem.FromModel(&t)

	if err != nil {
		return domain.NilID, err
	}

	result, err := repo.collection.InsertOne(ctx, bson.D{
		{"_id", todoItem.ID},
		{"title", todoItem.Title},
		{"description", todoItem.Description},
		{"priority_level", todoItem.Priority},
		{"completed", todoItem.Completed},
		{"created_at", todoItem.CreatedAt},
		{"updated_at", todoItem.UpdatedAt},
	})

	if err != nil {
		return domain.NilID, err
	}

	return domain.ID(result.InsertedID.(primitive.ObjectID).Hex()), nil
}

func (repo *TodoAdapter) Update(todo todo.Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(string(todo.ID))

	if err != nil {
		return err
	}

	//Filter
	filter := bson.M{"_id": bson.M{"$eq": oid}}

	//Update fields
	document := bson.M{"$set": bson.M{
		"title":       todo.Title,
		"description": todo.Description,
		"priority":    todo.Priority,
		"completed":   todo.Completed,
		"updated_at":  todo.UpdatedAt,
	}}

	//Update Item
	result, err := repo.collection.UpdateOne(ctx, filter, document)

	if err != nil {
		return err
	}

	if result.MatchedCount > 0 {
		return nil
	}

	return ErrNoItemsUpdated
}

func (repo *TodoAdapter) Delete(id domain.ID) error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(string(id))

	if err != nil {
		return err
	}

	filter := bson.M{"_id": bson.M{"$eq": oid}}

	result, err := repo.collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if result.DeletedCount > 0 {
		return nil
	}
	return ErrNoItemsDeleted
}
