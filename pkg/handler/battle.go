package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
	"net/http"
	"strconv"
)

// @Summary Create Battle
// @Security ApiKeyAuth
// @Tags battles
// @Description create battle
// @ID create-battle
// @Accept json
// @Produce json
// @Param input body carsBrandsBattleGame.Battle true "battle info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/battles [post]
func (h *Handler) createBattle(c *gin.Context) {
	fmt.Println("in")
	_, err := getUserId(c)
	if err != nil {
		return
	}

	var input carsBrandsBattle.Battle
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Battle.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllBattleResponse struct {
	Data []carsBrandsBattle.Battle `json:data`
}

// @Summary Get All Battles
// @Security ApiKeyAuth
// @Tags battles
// @Description get all battles
// @ID get-all-battles
// @Accept json
// @Produce json
// @Success 200 {object} getAllBattleResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/battles [get]
func (h *Handler) getAllBattles(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	battles, err := h.services.Battle.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBattleResponse{
		Data: battles,
	})
}

// @Summary Get Battle By Id
// @Security ApiKeyAuth
// @Tags battles
// @Description get battle by id
// @ID get-battle-by-id
// @Accept json
// @Produce json
// @Param id path int true "battle id"
// @Success 200 {object} carsBrandsBattleGame.Battle
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/battles/{id} [get]
func (h *Handler) getBattleById(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	battle, err := h.services.Battle.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, battle)
}

// @Summary Update Battle
// @Security ApiKeyAuth
// @Tags battles
// @Description update battle
// @ID update-battle
// @Accept json
// @Produce json
// @Param id path int true "battle id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/battles/{id} [put]
func (h *Handler) updateBattle(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input carsBrandsBattle.UpdateBattleInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Battle.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete Battle
// @Security ApiKeyAuth
// @Tags battles
// @Description delete battle
// @ID delete-battle
// @Accept json
// @Produce json
// @Param id path int true "battle id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/battles/{id} [delete]
func (h *Handler) deleteBattle(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.Battle.Delete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
