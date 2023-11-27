package handler

import (
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Error *string `json:"error" binding:"required"`
}

type statusResponse struct {
	Status string `json:"status" binding:"required"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{Error: &message})
}
