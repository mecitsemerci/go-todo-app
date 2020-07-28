package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mecitsemerci/clean-go-todo-api/app/api/controller"
	"github.com/mecitsemerci/clean-go-todo-api/app/api/dto"
	todoDto "github.com/mecitsemerci/clean-go-todo-api/app/api/dto/todo"
	"github.com/mecitsemerci/clean-go-todo-api/app/core/services"
	"net/http"
)

type TodoController struct {
	controller.ApiController
	TodoService services.TodoService
}

func (controller *TodoController) RegisterRoutes(apiRouteGroup *gin.RouterGroup) {
	router := apiRouteGroup.Group("/todo")
	{
		router.GET("/", controller.FindAll)
		router.GET("/:id", controller.Find)
		router.POST("/", controller.Create)
		router.PUT("/:id", controller.Update)
		router.DELETE("/:id", controller.Delete)
	}
}

// FindAll godoc
// @Summary Find all todo
// @Description Get all todo array
// @Tags Todo
// @Accept json
// @Produce json
// @Success 200 {array} todoDto.TodoOutput
// @Failure 400 {object} dto.ErrorOutput
// @Failure 422 {object} dto.ErrorOutput
// @Failure 500 {object} dto.ErrorOutput
// @Router /api/v1/todo [get]
func (controller *TodoController) FindAll(ctx *gin.Context) {
	todoList, err := controller.TodoService.GetAll()

	if err != nil {
		dto.NewErrorOutput(ctx, http.StatusUnprocessableEntity, "Something went wrong!", err)
		return
	}
	var result []todoDto.TodoOutput
	todoOutputDto := todoDto.TodoOutput{}
	for _, entity := range todoList {
		result = append(result, todoOutputDto.FromEntity(entity))
	}
	ctx.JSON(http.StatusOK, result)
}

// Find Todo godoc
// @Summary Find a todo
// @Description Get todo by id
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path string true "Todo Id"
// @Success 200 {object} todoDto.TodoOutput
// @Failure 400 {object} dto.ErrorOutput
// @Failure 404 {object} dto.ErrorOutput
// @Failure 500 {object} dto.ErrorOutput
// @Router /api/v1/todo/{id} [get]
func (controller *TodoController) Find(ctx *gin.Context) {

	todoId, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		dto.NewErrorOutput(ctx, http.StatusBadRequest, "Id is invalid.", err)
		return
	}

	todoEntity, err := controller.TodoService.Find(todoId)

	if err != nil {
		dto.NewErrorOutput(ctx, http.StatusNotFound, "Item is not exist.", err)
		return
	}

	todoOutput := todoDto.TodoOutput{}

	ctx.JSON(http.StatusOK, todoOutput.FromEntity(todoEntity))
}

// Create Todo godoc
// @Summary Create a todo
// @Description add by json todo
// @Tags Todo
// @Accept  json
// @Produce  json
// @Param account body todoDto.CreateTodoInput true "Create todo"
// @Success 200 {object} todoDto.CreateTodoOutput
// @Failure 400 {object} dto.ErrorOutput
// @Failure 422 {object} dto.ErrorOutput
// @Failure 500 {object} dto.ErrorOutput
// @Router /api/v1/todo/ [post]
func (controller *TodoController) Create(ctx *gin.Context) {
	var createTodoInput todoDto.CreateTodoInput
	if err := ctx.ShouldBindJSON(&createTodoInput); err != nil {
		dto.NewErrorOutput(ctx, http.StatusBadRequest, "Request model is invalid.", err)
		return
	}

	todoId, err := controller.TodoService.Create(createTodoInput.ToEntity())

	if err != nil {
		dto.NewErrorOutput(ctx, http.StatusUnprocessableEntity, "The item failed to create.", err)
		return
	}

	ctx.JSON(http.StatusCreated, todoDto.CreateTodoOutput{
		TodoId: todoId,
	})

}

// Update Todo godoc
// @Summary Update a todo
// @Description update by json todo
// @Tags Todo
// @Accept  json
// @Produce  json
// @Param id path string true "Todo Id"
// @Param account body todoDto.UpdateTodoInput true "Update todo"
// @Success 204 {object} todoDto.TodoOutput
// @Failure 400 {object} dto.ErrorOutput
// @Failure 404 {object} dto.ErrorOutput
// @Failure 500 {object} dto.ErrorOutput
// @Router /api/v1/todo/{id} [put]
func (controller *TodoController) Update(ctx *gin.Context) {

	todoId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		dto.NewErrorOutput(ctx, http.StatusBadRequest, "Id is invalid.", err)
		return
	}
	var updateTodoInput todoDto.UpdateTodoInput
	if err := ctx.ShouldBindJSON(&updateTodoInput); err != nil {
		dto.NewErrorOutput(ctx, http.StatusBadRequest, "Request model is invalid.", err)
		return
	}

	err = controller.TodoService.Update(updateTodoInput.ToEntity(todoId))

	if err != nil {
		dto.NewErrorOutput(ctx, http.StatusNotFound, "The item failed to update.", err)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})

}

// Delete Todo godoc
// @Summary Delete a todo
// @Description Delete by todo id
// @Tags Todo
// @Accept  json
// @Produce  json
// @Param id path string true "Todo Id"
// @Success 204 {object} todoDto.TodoOutput
// @Failure 400 {object} dto.ErrorOutput
// @Failure 404 {object} dto.ErrorOutput
// @Failure 500 {object} dto.ErrorOutput
// @Router /api/v1/todo/{id} [delete]
func (controller *TodoController) Delete(ctx *gin.Context) {

	todoId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		dto.NewErrorOutput(ctx, http.StatusBadRequest, "Id is invalid.", err)
		return
	}

	err = controller.TodoService.Delete(todoId)

	if err != nil {
		dto.NewErrorOutput(ctx, http.StatusNotFound, "The item failed to delete.", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
