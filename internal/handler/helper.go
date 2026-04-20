package handler

import (
	"github.com/dwikikf/alhikmah-api/internal/dto"
	"github.com/gin-gonic/gin"
)

func SuccessResponse[T any](c *gin.Context, code int, message string, data T) {
	c.JSON(code, dto.Response[T]{
		Status:  "Sukses",
		Message: message,
		Data:    data,
	})
}

func ErrorResponse[T any](c *gin.Context, code int, message string, errors T) {
	c.JSON(code, dto.Response[T]{
		Status:  "Error",
		Message: message,
		Errors:  errors,
	})
}
