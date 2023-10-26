package handler

import (
	"github.com/gin-gonic/gin"
	"name-service/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		human := api.Group("/human")
		{
			human.POST("/", h.createHuman)
			human.GET("/", h.getAllHumans)
			human.PUT("/:id", h.updateHuman)
			human.DELETE("/:id", h.deleteHuman)
		}
	}

	return router
}
