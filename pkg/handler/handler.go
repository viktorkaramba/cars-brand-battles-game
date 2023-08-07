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
		brands := api.Group("/brands")
		{
			brands.POST("/", h.createBrand)
			brands.GET("/", h.getAllBrands)
			brands.GET("/random", h.getBrandByRandom)
			brands.GET("/:id", h.getBrandById)
			brands.PUT("/:id", h.updateBrand)
			brands.DELETE("/:id", h.deleteBrand)
		}
		battles := api.Group("/battles")
		{
			battles.POST("/")
			battles.GET("/")
			battles.GET("/:id")
			battles.PUT("/:id")
			battles.DELETE("/:id")
		}
	}
	return router
}
