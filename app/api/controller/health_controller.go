package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mecitsemerci/clean-go-todo-api/app/api/dto"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/adapter/mongodb"
	"log"
	"net/http"
)

type HealthController struct{}

func (controller *HealthController) Init() *HealthController {
	return controller
}

func (controller *HealthController) RegisterRoutes(apiRouteGroup *gin.RouterGroup) {
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
	db := new(mongodb.DbContext)
	db.Connect()
	err := db.Client.Ping(db.Context, nil)
	defer db.Disconnect()
	mongo := map[string]bool{
		"mongodb": true,
	}
	log.Println("############# HealthyDependencies #############")
	var output = dto.HealthOutput{Status: "healthy", Dependencies: &mongo}
	if err != nil {
		output.Status = "unhealthy"
		mongo["mongodb"] = false
		output.Dependencies = &mongo
		fmt.Println(err.Error())
		ctx.JSON(http.StatusServiceUnavailable, output)
		return
	}
	ctx.JSON(http.StatusOK, output)
}
