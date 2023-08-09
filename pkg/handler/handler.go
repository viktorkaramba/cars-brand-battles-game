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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
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
			battles.POST("/", h.createBattle)
			battles.GET("/", h.getAllBattles)
			battles.GET("/:id", h.getBattleById)
			battles.PUT("/:id", h.updateBattle)
			battles.DELETE("/:id", h.deleteBattle)
		}
		scores := api.Group("/scores")
		{
			scores.GET("/", h.getAllScore)
			scores.GET("/:id", h.getScoreById)
			scores.PUT("/:id", h.updateScore)
			scores.DELETE("/:id", h.deleteScore)
		}
	}
	return router
}
