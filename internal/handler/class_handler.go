package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dwikikf/alhikmah-api/internal/dto"
	"github.com/dwikikf/alhikmah-api/internal/repository"
	"github.com/dwikikf/alhikmah-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ClassHandler struct {
	Usecase *usecase.ClassUsecase
}

func NewClassHandler(usecase *usecase.ClassUsecase) *ClassHandler {
	return &ClassHandler{Usecase: usecase}
}

func (h *ClassHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	classes, err := h.Usecase.GetAllClasses(ctx)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	if len(classes) == 0 {
		SuccessResponse[dto.ClassResponse](c, http.StatusOK, "no classes found", nil)
		return
	}

	SuccessResponse(c, http.StatusOK, "classes retrieved successfully", &classes)

}

func (h *ClassHandler) GetClassByID(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid class ID")
		return
	}

	class, err := h.Usecase.GetClassByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			ErrorResponse(c, http.StatusNotFound, "class not found")
			return
		}

		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	SuccessResponse(c, http.StatusOK, "class retrieved successfully", class)
}

func (h *ClassHandler) CreateClass(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.CreateClassRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	newClass := dto.CreateClassRequest{
		Code:  req.Code,
		Name:  req.Name,
		Grade: req.Grade,
	}

	id, err := h.Usecase.CreateClass(ctx, newClass)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicate) {
			ErrorResponse(c, http.StatusConflict, "class code already exists")
			return
		}

		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	SuccessResponse(c, http.StatusCreated, "class created successfully", &id)
}

func (h *ClassHandler) UpdateClass(c *gin.Context) {
	ctx := c.Request.Context()

	// validasi ID
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid class ID")
		return
	}

	var req dto.UpdateClassRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	updateClass := dto.UpdateClassRequest{
		Code:  req.Code,
		Name:  req.Name,
		Grade: req.Grade,
	}

	err = h.Usecase.UpdateClass(ctx, id, updateClass)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			ErrorResponse(c, http.StatusNotFound, "class not found")
			return
		}

		if errors.Is(err, repository.ErrDuplicate) {
			ErrorResponse(c, http.StatusConflict, "class code already exists")
			return
		}

		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	SuccessResponse[dto.ClassResponse](c, http.StatusOK, "class updated successfully", nil)

}

func (h *ClassHandler) DeleteClass(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid class ID")
		return
	}

	err = h.Usecase.DeleteClass(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			ErrorResponse(c, http.StatusNotFound, "class not found")
			return
		}

		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	SuccessResponse[dto.ClassResponse](c, http.StatusOK, "class deleted successfully", nil)
}
