package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/viktorkaramba/cars-brand-random-generator-app/pkg/service"
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
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createBrand)
			lists.GET("/", h.getAllBrands)
			lists.GET("/random", h.getBrandByRandom)
			lists.GET("/:id", h.getBrandById)
			lists.PUT("/:id", h.updateBrand)
			lists.DELETE("/:id", h.deleteBrand)
		}
	}
	return router
}
