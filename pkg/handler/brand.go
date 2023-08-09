package handler

import (
	"github.com/gin-gonic/gin"
	carsBrandRandomGenerator "github.com/viktorkaramba/cars-brand-random-generator-app"
	"net/http"
	"strconv"
)

func (h *Handler) createBrand(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	var input carsBrandRandomGenerator.Brand
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
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
	Data []carsBrandRandomGenerator.Brand `json:"data"`
}

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
	var input carsBrandRandomGenerator.UpdateBrandInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Brand.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

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
