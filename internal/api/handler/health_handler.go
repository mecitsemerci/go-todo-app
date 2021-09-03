package handler

import (
	"net/http"

	"github.com/opentracing/opentracing-go"

	"github.com/gin-gonic/gin"
	"github.com/mecitsemerci/go-todo-app/internal/api/dto"
)

//HealthHandler handles all health check operations
type HealthHandler struct{}

//Register maps HTTP operations with methods according to the router group
func (h *HealthHandler) Register(apiRouteGroup *gin.RouterGroup) {
	apiRouteGroup.GET("/status", h.Status)
}

// Status Health check godoc
// @Summary Check api status
// @Description Get api pulse status
// @Tags Health Check
// @Accept json
// @Produce json
// @Success 200 {object} dto.HealthOutput
// @Router /api/status [get]
func (h *HealthHandler) Status(ctx *gin.Context) {
	span := opentracing.GlobalTracer().StartSpan("HealthCheck")
	defer span.Finish()
	ctx.JSON(http.StatusOK, dto.HealthOutput{
		Status: "pass",
	})
}
