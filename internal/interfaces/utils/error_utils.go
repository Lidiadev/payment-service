package utils

import (
	"github.com/gin-gonic/gin"
	"payment-service/internal/interfaces/models"
)

func SendErrorResponse(c *gin.Context, statusCode int, errMsg string) {
	c.JSON(statusCode, models.ErrorResponse{
		Error: errMsg,
	})
}
