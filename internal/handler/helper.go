package handler

import (
	"github.com/dwikikf/alhikmah-api/internal/dto"
	"github.com/gin-gonic/gin"
)

func SuccessResponse[T any](c *gin.Context, code int, message string, data *T) {
	c.JSON(code, dto.Response[T]{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, dto.Response[any]{
		Status:  "error",
		Message: message,
	})
}
