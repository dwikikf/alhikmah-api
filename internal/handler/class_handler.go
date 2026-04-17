package handler

import (
	"github.com/dwikikf/alhikmah-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ClassHandler struct {
	Usecase *usecase.ClassUseCase
}

func NewClassHandler(usecase *usecase.ClassUseCase) *ClassHandler {
	return &ClassHandler{Usecase: usecase}
}

func (h *ClassHandler) GetAllClasses(c *gin.Context) {

}

func (h *ClassHandler) GetClassByID(c *gin.Context) {

}

func (h *ClassHandler) CreateClass(c *gin.Context) {
	// Implementation for creating a new class
}

func (h *ClassHandler) UpdateClass(c *gin.Context) {
	// Implementation for updating an existing class
}

func (h *ClassHandler) DeleteClass(c *gin.Context) {
	// Implementation for deleting a class
}
