package handler

import (
	"errors"
	"net/http"

	"github.com/dwikikf/alhikmah-api/internal/dto"
	"github.com/dwikikf/alhikmah-api/internal/repository"
	"github.com/dwikikf/alhikmah-api/internal/usecase"
	"github.com/dwikikf/alhikmah-api/pkg/validator"
	"github.com/gin-gonic/gin"
)

type AttendanceHandler struct {
	Usecase   *usecase.AttendanceUsecaseImpl
	Validator *validator.CustomValidator
}

func NewAttendanceHandler(usecase *usecase.AttendanceUsecaseImpl, validator *validator.CustomValidator) *AttendanceHandler {
	return &AttendanceHandler{Usecase: usecase, Validator: validator}
}

func (h *AttendanceHandler) CheckedIn(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.CreateAttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid request body", nil)
		return
	}

	if validationErr := h.Validator.Validate(req); validationErr != nil {
		ErrorResponse(c, http.StatusUnprocessableEntity, "validation error", &validationErr)
		return
	}

	id, err := h.Usecase.CheckIn(ctx, req)
	if err != nil {
		if errors.Is(err, repository.ErrAlreadyCheckedIn) {
			ErrorResponse(c, http.StatusConflict, "student already checked in", nil)
			return
		}
		ErrorResponse(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	SuccessResponse(c, http.StatusCreated, "Terima Kasih", &id)
}
