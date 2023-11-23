package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "github.com/viktorkaramba/cars-brand-random-generator-app/docs"
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
		auth.POST("logout", h.logout)
	}

	api := router.Group("/api", h.userIdentity)
	{

		users := api.Group("/users")
		{
			users.GET("/:username", h.getUserByUsername)
		}
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
		userStatistics := api.Group("/user-statistics")
		{
			userStatistics.GET("/by-score", h.getGeneralUserStatisticsByScore)
		}
		userInterfaceData := api.Group("/user-interface-data")
		{
			userInterfaceData.GET("/", h.getGeneralUserInterfaceData)
			userInterfaceData.GET("/:id", h.getGeneralUserInterfaceDataByBattleId)
		}
		userHistory := api.Group("/users-history")
		{
			userHistory.GET("/", h.getAllUsersHistory)
			userHistory.GET("/:id", h.getUsersHistoryByBattleId)
		}
	}
	router.POST("/refresh-token", h.refreshToken)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
