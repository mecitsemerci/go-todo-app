package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mecitsemerci/go-todo-app/internal/api/dto"
	"net/http"
)

type HealthController struct {
}

func NewHealthController() HealthController {
	return HealthController{}
}

func (controller *HealthController) Register(apiRouteGroup *gin.RouterGroup) {
	apiRouteGroup.GET("/status", controller.Status)
}

// Health check godoc
// @Summary Check api status
// @Description Get api healthy status
// @Tags Health Check
// @Accept json
// @Produce json
// @Success 200 {object} dto.HealthOutput
// @Router /api/status [get]
func (controller *HealthController) Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, dto.HealthOutput{
		Status: "ok",
	})
}

