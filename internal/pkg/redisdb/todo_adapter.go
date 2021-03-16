package redisdb

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain/todo"
	"github.com/opentracing/opentracing-go"
)

const (
	//CollectionName is Redis Key items collection name
	CollectionName = "todos"
)

//TodoAdapter is a redis implementation of TodoRepository
type TodoAdapter struct {
	client *redis.Client
}

//NewTodoAdapter returns a redis adapter according to TodoRepository
func NewTodoAdapter(client *redis.Client) *TodoAdapter {
	return &TodoAdapter{
		client: client,
	}
}

// GetAll returns all items
func (a *TodoAdapter) GetAll(ctx context.Context) ([]todo.Todo, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "TodoAdapter-GetAll")
	defer span.Finish()

	items, err := a.client.HVals(ctx, CollectionName).Result()

	var todos = make([]todo.Todo, 0, len(items))

	if err != nil {
		return todos, err
	}

	var errs = make([]string, 0)
	var tmpTodo *todo.Todo
	for _, item := range items {
		err = json.Unmarshal([]byte(item), &tmpTodo)
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		todos = append(todos, *tmpTodo)
	}

	if len(errs) == 0 {
		return todos, nil
	}

	return todos, errors.New(strings.Join(errs, ";"))
}

//GetByID retrieve item according to given item ID
func (a *TodoAdapter) GetByID(ctx context.Context, id domain.ID) (todo.Todo, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "TodoAdapter-GetByID")
	defer span.Finish()

	var t todo.Todo

	byt, err := a.client.HGet(ctx, CollectionName, id.String()).Bytes()

	if err != nil {
		return todo.DefaultTodo, err
	}

	err = json.Unmarshal(byt, &t)

	if err != nil {
		return todo.DefaultTodo, err
	}
	return t, nil
}

//Insert item according to given item
func (a *TodoAdapter) Insert(ctx context.Context, t todo.Todo) (domain.ID, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "TodoAdapter-Insert")
	defer span.Finish()

	serializedValue, err := json.Marshal(t)
	if err != nil {
		return domain.ZeroID, err
	}
	err = a.client.HSet(ctx, CollectionName, map[string]interface{}{
		t.ID.String(): serializedValue,
	}).Err()

	if err != nil {
		return domain.ZeroID, err
	}

	return t.ID, nil
}

//Update item according to given item
func (a *TodoAdapter) Update(ctx context.Context, t todo.Todo) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "TodoAdapter-Update")
	defer span.Finish()

	//Set created date according to existing item
	existingTodo, err := a.GetByID(ctx, t.ID)

	if err == nil {
		t.CreatedAt = existingTodo.CreatedAt
	}

	serializedValue, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return a.client.HSet(ctx, CollectionName, map[string]interface{}{
		t.ID.String(): serializedValue,
	}).Err()
}

//Delete item according to given todoID
func (a *TodoAdapter) Delete(ctx context.Context, id domain.ID) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "TodoAdapter-Delete")
	defer span.Finish()

	return a.client.HDel(ctx, CollectionName, id.String()).Err()
}

//Close disconnects the connection
func (a *TodoAdapter) Close(ctx context.Context) error {
	return a.client.Close()
}
