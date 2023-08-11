package handler

import (
	"github.com/gin-gonic/gin"
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
	"net/http"
)

type getGeneralUserStatisticsByScoreResponse struct {
	Data []carsBrandsBattle.UserStatistics `json:data`
}

// @Summary Get General User Statistics By Score
// @Security ApiKeyAuth
// @Tags userStatistics
// @Description get user statistics by score
// @ID get-user-statistics
// @Accept json
// @Produce json
// @Success 200 {object} getGeneralUserStatisticsByScoreResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user-statistics/by-score [get]
func (h *Handler) getGeneralUserStatisticsByScore(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	userStatistics, err := h.services.UserStatistics.GetGeneralStatisticsByScore()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getGeneralUserStatisticsByScoreResponse{
		Data: userStatistics,
	})
}
