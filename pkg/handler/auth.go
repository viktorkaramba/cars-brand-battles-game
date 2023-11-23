package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
	"net/http"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept json
// @Produce json
// @Param input body carsBrandsBattleGame.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	fmt.Println("in")
	var input carsBrandsBattle.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	_, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	user, token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newToken := carsBrandsBattle.Token{TokenValue: token, Revoked: false, UserId: user.Id}
	_, err = h.services.Tokens.Create(newToken)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token":    token,
		"userId":   user.Id,
		"name":     user.Name,
		"username": user.Username,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept json
// @Produce json
// @Param input body signInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newToken := carsBrandsBattle.Token{TokenValue: token, Revoked: false, UserId: user.Id}
	_, err = h.services.Tokens.Create(newToken)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token":    token,
		"userId":   user.Id,
		"name":     user.Name,
		"username": user.Username,
	})
}

func (h *Handler) logout(c *gin.Context) {
	token, err := checkHeaderToken(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	revokedToken := carsBrandsBattle.UpdateTokenInput{
		Revoked: func() *bool { b := true; return &b }(),
	}
	err = h.services.Tokens.Update(token, revokedToken)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) getUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := h.services.Authorization.GetUserByUsername(username)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": user.Id,
	})
}

func (h *Handler) refreshToken(c *gin.Context) {

	var input carsBrandsBattle.RefreshTokenInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	newToken, err := h.services.RefreshToken(input.UserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken": newToken,
	})
}
