package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Get All User History
// @Security ApiKeyAuth
// @Tags userHistory
// @Description get users history
// @ID get-user-history
// @Accept json
// @Produce json
// @Success 200 {object} UserInterfaceData
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/users-history [get]
func (h *Handler) getAllUsersHistory(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	userInterfaceData, err := h.services.UserInterfaceData.GetAll(true)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, userInterfaceData)
}

// @Summary Get User History By BattleId
// @Security ApiKeyAuth
// @Tags userHistoryByBattleID
// @Description get users history by battleId
// @ID get-user-history-by-battle-id
// @Accept json
// @Produce json
// @Success 200 {object} UserInterfaceData
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/users-history/{id} [get]
func (h *Handler) getUsersHistoryByBattleId(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	userInterfaceData, err := h.services.UserInterfaceData.GetById(id, true)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, userInterfaceData)
}
