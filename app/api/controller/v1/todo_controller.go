package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mecitsemerci/clean-go-todo-api/app/api/controller"
	todoDto "github.com/mecitsemerci/clean-go-todo-api/app/api/dto/todo"
	"github.com/mecitsemerci/clean-go-todo-api/app/core/services"
	"github.com/swaggo/swag/example/celler/httputil"
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
// @Failure 400 {object} httputil.HTTPError
// @Failure 422 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/todo [get]
func (controller *TodoController) FindAll(ctx *gin.Context) {
	if todoList, err := controller.TodoService.GetAll(); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	} else {
		var result []todoDto.TodoOutput
		todoOutputDto := todoDto.TodoOutput{}
		for _, entity := range todoList {
			result = append(result, todoOutputDto.FromEntity(entity))
		}
		ctx.JSON(http.StatusOK, result)
	}
}

// Find Todo godoc
// @Summary Find a todo
// @Description Get todo by id
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path string true "Todo Id"
// @Success 200 {object} todoDto.TodoOutput
// @Failure 400 {object} httputil.HTTPError
// @Failure 422 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/todo/{id} [get]
func (controller *TodoController) Find(ctx *gin.Context) {

	todoId, err := uuid.Parse(ctx.Param("id"))

	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	todoEntity, err := controller.TodoService.Find(todoId)

	if err != nil {
		httputil.NewError(ctx, http.StatusUnprocessableEntity, err)
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
// @Failure 400 {object} httputil.HTTPError
// @Failure 422 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/todo/ [post]
func (controller *TodoController) Create(ctx *gin.Context) {
	var createTodoInput todoDto.CreateTodoInput
	if err := ctx.ShouldBindJSON(&createTodoInput); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	todoId, err := controller.TodoService.Create(createTodoInput.ToEntity())

	if err != nil {
		httputil.NewError(ctx, http.StatusUnprocessableEntity, err)
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
// @Success 200 {object} todoDto.TodoOutput
// @Failure 400 {object} httputil.HTTPError
// @Failure 422 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/todo/{id} [put]
func (controller *TodoController) Update(ctx *gin.Context) {

	todoId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	var updateTodoInput todoDto.UpdateTodoInput
	if err := ctx.ShouldBindJSON(&updateTodoInput); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	isUpdated, err := controller.TodoService.Update(updateTodoInput.ToEntity(todoId))

	if err != nil || !isUpdated {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})

}

// Delete Todo godoc
// @Summary Delete a todo
// @Description Delete by todo id
// @Tags Todo
// @Accept  json
// @Produce  json
// @Param id path string true "Todo Id"
// @Success 204 {object} todoDto.TodoOutput
// @Failure 400 {object} httputil.HTTPError
// @Failure 422 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /api/v1/todo/{id} [delete]
func (controller *TodoController) Delete(ctx *gin.Context) {

	todoId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	isDeleted, err := controller.TodoService.Delete(todoId)

	if err != nil || !isDeleted {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{})

}
