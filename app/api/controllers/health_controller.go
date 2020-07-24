package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mecitsemerci/clean-go-todo-api/app/api/dtos"
	"net/http"
)

type HealthController struct {
	ApiController
}

func (controller *HealthController) RegisterRoutes(apiRouteGroup *gin.RouterGroup) {
	router := apiRouteGroup.Group("/health")
	router.GET("/", controller.Healthy)

}

func (controller *HealthController) Healthy(ctx *gin.Context) {
	status := dtos.HealthOutputDto{
		Status: "healthy",
	}
	ctx.JSON(http.StatusOK, status)
}
