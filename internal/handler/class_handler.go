package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dwikikf/alhikmah-api/internal/dto"
	"github.com/dwikikf/alhikmah-api/internal/repository"
	"github.com/dwikikf/alhikmah-api/internal/usecase"
	"github.com/dwikikf/alhikmah-api/pkg/validator"
	"github.com/gin-gonic/gin"
)

type ClassHandler struct {
	Usecase   *usecase.ClassUsecase
	validator *validator.CustomValidator
}

func NewClassHandler(usecase *usecase.ClassUsecase, validator *validator.CustomValidator) *ClassHandler {
	return &ClassHandler{
		Usecase:   usecase,
		validator: validator,
	}
}

func (h *ClassHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	classes, err := h.Usecase.GetAllClasses(ctx)
	if err != nil {
		ErrorResponse[any](c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	if len(classes) == 0 {
		SuccessResponse[any](c, http.StatusOK, "no classes found", nil)
		return
	}

	SuccessResponse(c, http.StatusOK, "classes retrieved successfully", &classes)

}

func (h *ClassHandler) GetClassByID(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ErrorResponse[any](c, http.StatusBadRequest, "invalid class ID", nil)
		return
	}

	class, err := h.Usecase.GetClassByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrClassNotFound) {
			ErrorResponse[any](c, http.StatusNotFound, "class not found", nil)
			return
		}

		ErrorResponse[any](c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	SuccessResponse(c, http.StatusOK, "class retrieved successfully", class)
}

func (h *ClassHandler) CreateClass(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.CreateClassRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse[any](c, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	id, err := h.Usecase.CreateClass(ctx, req)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicate) {
			ErrorResponse[any](c, http.StatusConflict, "class code already exists", nil)
			return
		}

		ErrorResponse[any](c, http.StatusInternalServerError, "internal server error", nil)
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
		ErrorResponse[any](c, http.StatusBadRequest, "invalid class ID", nil)
		return
	}

	var req dto.UpdateClassRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse[any](c, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	err = h.Usecase.UpdateClass(ctx, id, req)
	if err != nil {
		if errors.Is(err, repository.ErrClassNotFound) {
			ErrorResponse[any](c, http.StatusNotFound, "class not found", nil)
			return
		}

		if errors.Is(err, repository.ErrDuplicate) {
			ErrorResponse[any](c, http.StatusConflict, "class code already exists", nil)
			return
		}

		ErrorResponse[any](c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	SuccessResponse[any](c, http.StatusOK, "class updated successfully", nil)

}

func (h *ClassHandler) DeleteClass(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ErrorResponse[any](c, http.StatusBadRequest, "invalid class ID", nil)
		return
	}

	err = h.Usecase.DeleteClass(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrClassNotFound) {
			ErrorResponse[any](c, http.StatusNotFound, "class not found", nil)
			return
		}

		if errors.Is(err, repository.ErrForeignKey) {
			ErrorResponse[any](c, http.StatusConflict, "cannot delete student with existing references", nil)
			return
		}

		ErrorResponse[any](c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	SuccessResponse[any](c, http.StatusOK, "class deleted successfully", nil)
}
