package mementodb

import (
	"github.com/google/uuid"
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/utils"
	"log"
)

var todoCollection = map[string]todo.Todo{}

type DbContext struct{}

func (db *DbContext) Seed() {
	now := utils.UtcNow()
	id1, err := uuid.Parse("2d554d6a-d908-4de8-929e-e9c4d487c6a0")
	log.Println(err)
	todoCollection["2d554d6a-d908-4de8-929e-e9c4d487c6a0"] = todo.Todo{
		Id:        id1,
		Title:     "Todo1",
		IsDone:    false,
		CreatedAt: now,
		UpdatedAt: now,
	}
	id2, _ := uuid.Parse("6ba7b811-9dad-11d1-80b4-00c04fd430c8")
	todoCollection["6ba7b811-9dad-11d1-80b4-00c04fd430c8"] = todo.Todo{
		Id:        id2,
		Title:     "Todo2",
		IsDone:    true,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
