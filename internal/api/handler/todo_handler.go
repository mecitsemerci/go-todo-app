package handler

import (
	"github.com/pkg/errors"
	"net/http"

	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/mecitsemerci/go-todo-app/internal/api/dto"
	"github.com/mecitsemerci/go-todo-app/internal/api/httperror"
	"github.com/mecitsemerci/go-todo-app/internal/core/domain"
	"github.com/mecitsemerci/go-todo-app/internal/core/interfaces"
	"github.com/mecitsemerci/go-todo-app/internal/pkg/check"
	"github.com/mecitsemerci/go-todo-app/internal/pkg/validator"
)

//TodoHandler handles all todo operations
type TodoHandler struct {
	TodoService interfaces.TodoService
}

//NewTodoHandler returns new TodoHandler instance
func NewTodoHandler(todoService interfaces.TodoService) TodoHandler {
	return TodoHandler{TodoService: todoService}
}

//Register maps HTTP operations with methods according to the router group
func (h *TodoHandler) Register(apiRouteGroup *gin.RouterGroup) {
	router := apiRouteGroup.Group("/todo")
	{
		router.GET("", h.GetAll)
		router.GET("/:id", h.Find)
		router.POST("", h.Create)
		router.PUT("/:id", h.Update)
		router.DELETE("/:id", h.Delete)
	}
}

// GetAll godoc
// @Summary Get all todo
// @Description Get all todo array
// @Tags Todo
// @Accept json
// @Produce json
// @Success 200 {array} dto.TodoOutput
// @Failure 400 {object} dto.ErrorOutput
// @Failure 422 {object} dto.ErrorOutput
// @Router /api/v1/todo [get]
func (h *TodoHandler) GetAll(ctx *gin.Context) {
	span := opentracing.GlobalTracer().StartSpan("TodoHandler-GetAll")
	spanContext := opentracing.ContextWithSpan(ctx.Request.Context(), span)
	defer span.Finish()

	todoList, err := h.TodoService.GetAll(spanContext)
	if err != nil {
		log.Error(errors.Wrap(err, "get all failed"))
		httperror.NewError(ctx, http.StatusUnprocessableEntity, "Something went wrong!", err)
		return
	}
	var result = make([]dto.TodoOutput, 0)
	todoOutputDto := dto.TodoOutput{}
	for _, model := range todoList {
		result = append(result, todoOutputDto.FromModel(model))
	}
	ctx.JSON(http.StatusOK, result)
}

// Find Todo godoc
// @Summary Find a todo
// @Description Get todo by id
// @Tags Todo
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} dto.TodoOutput
// @Failure 400 {object} dto.ErrorOutput
// @Failure 404 {object} dto.ErrorOutput
// @Router /api/v1/todo/{id} [get]
func (h *TodoHandler) Find(ctx *gin.Context) {
	span := opentracing.GlobalTracer().StartSpan("TodoHandler-Find")
	spanContext := opentracing.ContextWithSpan(ctx.Request.Context(), span)
	defer span.Finish()

	todoID := ctx.Param("id")
	if check.IsEmptyOrWhiteSpace(todoID) {
		log.Error("todo ID is empty error")
		httperror.NewError(ctx, http.StatusBadRequest, "ID is empty.", nil)
		return
	}

	todoModel, err := h.TodoService.Find(spanContext, domain.ID(todoID))
	if err != nil {
		log.Error(errors.Wrap(err, "item not exist"))
		httperror.NewError(ctx, http.StatusUnprocessableEntity, "Item is not exist.", err)
		return
	}
	todoOutput := dto.TodoOutput{}

	ctx.JSON(http.StatusOK, todoOutput.FromModel(todoModel))
}

// Create Todo godoc
// @Summary Create a todo
// @Description add by json todo
// @Tags Todo
// @Accept  json
// @Produce  json
// @Param todo body dto.CreateTodoInput true "Create todo"
// @Success 200 {object} dto.CreateTodoOutput
// @Failure 400 {object} dto.ErrorOutput
// @Failure 422 {object} dto.ErrorOutput
// @Router /api/v1/todo [post]
func (h *TodoHandler) Create(ctx *gin.Context) {
	span := opentracing.GlobalTracer().StartSpan("TodoHandler-Create")
	spanContext := opentracing.ContextWithSpan(ctx.Request.Context(), span)
	defer span.Finish()

	var createTodoInput dto.CreateTodoInput

	if err := ctx.ShouldBindJSON(&createTodoInput); err != nil {
		log.Error(errors.Wrap(err, "invalid request error"))
		httperror.NewError(ctx, http.StatusBadRequest, "Request model is invalid.", err)
		return
	}

	if err := validator.Validate(createTodoInput); err != nil {
		log.Error(errors.Wrap(err, "validation error"))
		httperror.NewError(ctx, http.StatusBadRequest, "Validation error", err)
		return
	}

	todoID, err := h.TodoService.Create(spanContext, createTodoInput.ToModel())

	if err != nil {
		log.Error(errors.Wrap(err, "item create error"))
		httperror.NewError(ctx, http.StatusUnprocessableEntity, "The item failed to create.", err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateTodoOutput{
		TodoID: todoID.String(),
	})

}

// Update Todo godoc
// @Summary Update a todo
// @Description update by json todo
// @Tags Todo
// @Accept  json
// @Produce  json
// @Param id path string true "Todo ID"
// @Param todo body dto.UpdateTodoInput true "Update todo"
// @Success 204 {object} dto.TodoOutput
// @Failure 400 {object} dto.ErrorOutput
// @Failure 404 {object} dto.ErrorOutput
// @Router /api/v1/todo/{id} [put]
func (h *TodoHandler) Update(ctx *gin.Context) {
	span := opentracing.GlobalTracer().StartSpan("TodoHandler-Create")
	spanContext := opentracing.ContextWithSpan(ctx.Request.Context(), span)
	defer span.Finish()

	todoID := ctx.Param("id")

	if check.IsEmptyOrWhiteSpace(todoID) {
		log.Error("todo ID is empty error")
		httperror.NewError(ctx, http.StatusBadRequest, "ID is invalid.", nil)
		return
	}

	var updateTodoInput dto.UpdateTodoInput

	if err := ctx.ShouldBindJSON(&updateTodoInput); err != nil {
		log.Error(errors.Wrap(err, "invalid request error"))
		httperror.NewError(ctx, http.StatusBadRequest, "Request model is invalid.", err)
		return
	}

	if err := validator.Validate(updateTodoInput); err != nil {
		log.Error(errors.Wrap(err, "validation error"))
		httperror.NewError(ctx, http.StatusBadRequest, "Validation error", err)
		return
	}

	model := updateTodoInput.ToModel(todoID)

	err := h.TodoService.Update(spanContext, model)

	if err != nil {
		log.Error(errors.Wrap(err, "item update error"))
		httperror.NewError(ctx, http.StatusNotFound, "The item failed to update.", err)
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
// @Param id path string true "Todo ID"
// @Success 204 {object} dto.TodoOutput
// @Failure 400 {object} dto.ErrorOutput
// @Failure 404 {object} dto.ErrorOutput
// @Router /api/v1/todo/{id} [delete]
func (h *TodoHandler) Delete(ctx *gin.Context) {
	span := opentracing.GlobalTracer().StartSpan("TodoHandler-Create")
	spanContext := opentracing.ContextWithSpan(ctx.Request.Context(), span)
	defer span.Finish()

	todoID := ctx.Param("id")

	if check.IsEmptyOrWhiteSpace(todoID) {
		log.Error("todo ID is empty error")
		httperror.NewError(ctx, http.StatusBadRequest, "ID is empty.", nil)
		return
	}

	err := h.TodoService.Delete(spanContext, domain.ID(todoID))

	if err != nil {
		log.Error(errors.Wrap(err, "item delete error"))
		httperror.NewError(ctx, http.StatusNotFound, "The item failed to delete.", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
