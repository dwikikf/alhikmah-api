package handler

import (
	"github.com/dwikikf/alhikmah-api/internal/dto"
	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, code int, message string, data any) {
	c.JSON(code, dto.Response{
		Status:  "Sukses",
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string, errors any) {
	c.JSON(code, dto.Response{
		Status:  "Error",
		Message: message,
		Errors:  errors,
	})
}
