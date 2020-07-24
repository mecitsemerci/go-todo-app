package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mecitsemerci/clean-go-todo-api/app/api/dtos"
	"github.com/mecitsemerci/clean-go-todo-api/app/core/services"
	"net/http"
)

type TodoController struct {
	ApiController
	TodoService services.TodoService
}

func (controller *TodoController) RegisterRoutes(apiRouteGroup *gin.RouterGroup) {
	router := apiRouteGroup.Group("/v1/todo")
	router.GET("/", controller.FindAll)
	router.GET("/:id", controller.Find)
	router.POST("/", controller.Create)
	router.PUT("/:id", controller.Update)
	router.DELETE("/:id", controller.Delete)
}

func (controller *TodoController) FindAll(ctx *gin.Context) {
	if todoList, err := controller.TodoService.GetAll(); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	} else {
		var result []dtos.TodoOutputDto
		todoOutputDto := dtos.TodoOutputDto{}
		for _, entity := range todoList {
			result = append(result, todoOutputDto.FromEntity(entity))
		}
		ctx.JSON(http.StatusOK, result)
	}
}

func (controller *TodoController) Find(ctx *gin.Context) {

	todoId, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "must provide a valid todo id")
		return
	}

	todoEntity, err := controller.TodoService.Find(todoId)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	todoOutput := dtos.TodoOutputDto{}

	ctx.JSON(http.StatusOK, todoOutput.FromEntity(todoEntity))

}
func (controller *TodoController) Create(ctx *gin.Context) {
	var createTodoInput dtos.CreateTodoInputDto
	if err := ctx.ShouldBindJSON(&createTodoInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todoId, err := controller.TodoService.Create(createTodoInput.ToEntity())

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"todo_id": todoId})

}
func (controller *TodoController) Update(ctx *gin.Context) {

	todoId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "must provide a valid todo id")
		return
	}
	var updateTodoInput dtos.UpdateTodoInputDto
	if err := ctx.ShouldBindJSON(&updateTodoInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}


	isUpdated, err := controller.TodoService.Update(updateTodoInput.ToEntity(todoId))

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := gin.H{"result": isUpdated}
	if !isUpdated {
		ctx.JSON(http.StatusNoContent, result)
	}
	ctx.JSON(http.StatusOK, result)

}
func (controller *TodoController) Delete(ctx *gin.Context) {

	todoId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "must provide a valid todo id")
		return
	}


	isDeleted, err := controller.TodoService.Delete(todoId)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := gin.H{"result": isDeleted}
	if !isDeleted {
		ctx.JSON(http.StatusNoContent, result)
	}
	ctx.JSON(http.StatusOK, result)

}
