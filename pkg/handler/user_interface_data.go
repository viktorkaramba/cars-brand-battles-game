package handler

import (
	"github.com/gin-gonic/gin"
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
	"net/http"
)

type getGeneralUserInterfaceDataResponse struct {
	Data []carsBrandsBattle.UserInterfaceData `json:data`
}

// @Summary Get General User Interface Data
// @Security ApiKeyAuth
// @Tags userInterfaceData
// @Description get user interface data
// @ID get-user-interface-data
// @Accept json
// @Produce json
// @Success 200 {object} getGeneralUserInterfaceDataResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user-interface-data [get]
func (h *Handler) getGeneralUserInterfaceData(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	userInterfaceData, err := h.services.UserInterfaceData.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getGeneralUserInterfaceDataResponse{
		Data: userInterfaceData,
	})
}
