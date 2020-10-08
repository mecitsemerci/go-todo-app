package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mecitsemerci/clean-go-todo-api/app/api/dto"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/adapter/mongodb"
	"net/http"
	"time"
)

type HealthController struct {
	DbContext mongodb.DbContext
}

func NewHealthController(dbContext mongodb.DbContext) HealthController {
	return HealthController{DbContext: dbContext}
}

func (controller *HealthController) Register(apiRouteGroup *gin.RouterGroup) {
	router := apiRouteGroup.Group("/health")
	router.GET("/", controller.Healthy)
	router.GET("/dependencies", controller.HealthyDependencies)

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

// Health check godoc
// @Summary Check api status
// @Description Get api healthy status with dependencies
// @Tags Health Check
// @Accept json
// @Produce json
// @Success 200 {object} dto.HealthOutput
// @Router /api/health/dependencies [get]
func (controller *HealthController) HealthyDependencies(ctx *gin.Context) {

	controller.DbContext.ConnectWithTimeout(1 * time.Second)

	err := controller.DbContext.Client.Ping(controller.DbContext.Context, nil)

	defer controller.DbContext.Disconnect()

	mongo := map[string]bool{"mongodb": true}

	var output = dto.HealthOutput{Status: "healthy", Dependencies: &mongo}

	if err != nil {
		output.Status = "unhealthy"
		mongo["mongodb"] = false
		output.Dependencies = &mongo
		ctx.JSON(http.StatusServiceUnavailable, output)
		return
	}

	ctx.JSON(http.StatusOK, output)
}
