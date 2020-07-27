package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mecitsemerci/clean-go-todo-api/app/api/dto"
	"net/http"
)

type HealthController struct {
	ApiController
}

func (controller *HealthController) RegisterRoutes(apiRouteGroup *gin.RouterGroup) {
	router := apiRouteGroup.Group("/health")
	router.GET("/", controller.Healthy)

}
// Health check godoc
// @Summary Check api status
// @Description Get api healthy status
// @Tags Health Check
// @Accept json
// @Produce json
// @Success 200 {object} dto.HealthOutput
// @Router /api/health [get]
func (controller *HealthController) Healthy(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, dto.HealthOutput{
		Status: "healthy",
	})
}
