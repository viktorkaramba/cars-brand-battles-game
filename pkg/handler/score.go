package handler

import (
	"github.com/gin-gonic/gin"
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
	"net/http"
	"strconv"
)

type getAllScoreResponse struct {
	Data []carsBrandsBattle.Score `json:data`
}

// @Summary Get All Scores
// @Security ApiKeyAuth
// @Tags scores
// @Description get all scores
// @ID get-all-scores
// @Accept json
// @Produce json
// @Success 200 {object} getAllScoreResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/scores [get]
func (h *Handler) getAllScore(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	scores, err := h.services.Score.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllScoreResponse{
		Data: scores,
	})
}

// @Summary Get Score By Id
// @Security ApiKeyAuth
// @Tags scores
// @Description get score by id
// @ID get-score-by-id
// @Accept json
// @Produce json
// @Param id path int true "score id"
// @Success 200 {object} carsBrandsBattle.Score
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/scores/{id} [get]
func (h *Handler) getScoreById(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	score, err := h.services.Score.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, score)
}

// @Summary Update Score
// @Security ApiKeyAuth
// @Tags scores
// @Description update score
// @ID update-score
// @Accept json
// @Produce json
// @Param id path int true "score id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/scores/{id} [put]
func (h *Handler) updateScore(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	var input carsBrandsBattle.UpdateScoreInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Score.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete Score
// @Security ApiKeyAuth
// @Tags scores
// @Description delete score
// @ID delete-score
// @Accept json
// @Produce json
// @Param id path int true "score id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/scores/{id} [delete]
func (h *Handler) deleteScore(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.Score.Delete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
