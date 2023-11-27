package handler

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
	"io"
	"net/http"
	"strconv"
)

// @Summary Create Brand
// @Security ApiKeyAuth
// @Tags brands
// @Description create brand
// @ID create-brand
// @Accept json
// @Produce json
// @Param input body carsBrandsBattleGame.Brand true "brand info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/brands [post]
func (h *Handler) createBrand(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	var input carsBrandsBattle.Brand
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Brand.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllBrandResponse struct {
	Data []carsBrandsBattle.Brand `json:"data"`
}

// @Summary Get All Brands
// @Security ApiKeyAuth
// @Tags brands
// @Description get all brands
// @ID get-all-brands
// @Accept json
// @Produce json
// @Success 200 {object} getAllBrandResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/brands [get]
func (h *Handler) getAllBrands(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	brands, err := h.services.Brand.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBrandResponse{
		Data: brands,
	})
}

// @Summary Get Brand By Id
// @Security ApiKeyAuth
// @Tags brands
// @Description get brand by id
// @ID get-brand-by-id
// @Accept json
// @Produce json
// @Param id path int true "brand id"
// @Success 200 {object} carsBrandsBattleGame.Brand
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/brands/{id} [get]
func (h *Handler) getBrandById(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	brand, err := h.services.Brand.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, brand)
}

// @Summary Get Random Brand
// @Security ApiKeyAuth
// @Tags brands
// @Description get random brand
// @ID get-random-brand
// @Accept json
// @Produce json
// @Success 200 {object} carsBrandsBattleGame.Brand
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/brands/random [get]
func (h *Handler) getBrandByRandom(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}
	brand, err := h.services.Brand.GetOneByRandom()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, brand)
}

// @Summary Update Brand
// @Security ApiKeyAuth
// @Tags brands
// @Description update brand
// @ID update-brand
// @Accept json
// @Produce json
// @Param id path int true "brand id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/brands/{id} [put]
func (h *Handler) updateBrand(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	// Check if there are any additional fields in the JSON body
	if err := h.validateJSONTags(body, carsBrandsBattle.UpdateBrandInput{}); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input := carsBrandsBattle.UpdateBrandInput{}
	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Brand.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete Brand
// @Security ApiKeyAuth
// @Tags brands
// @Description delete brand
// @ID delete-brand
// @Accept json
// @Produce json
// @Param id path int true "brand id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/brands/{id} [delete]
func (h *Handler) deleteBrand(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.Brand.Delete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
