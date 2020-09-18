package adapter

import (
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/adapter/mongodb"
)

//DBContext
func ProvideDbContext() mongodb.DbContext  {
	return mongodb.NewDbContext()
}

//Repository
func ProvideTodoRepository(dbContext mongodb.DbContext) todo.ITodoRepository {
	return mongodb.NewTodoAdapter(dbContext)
}
