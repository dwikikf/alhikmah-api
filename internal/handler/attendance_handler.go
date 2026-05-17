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

type AttendanceHandler struct {
	Usecase   *usecase.AttendanceUsecaseImpl
	Validator *validator.CustomValidator
}

func NewAttendanceHandler(usecase *usecase.AttendanceUsecaseImpl, validator *validator.CustomValidator) *AttendanceHandler {
	return &AttendanceHandler{Usecase: usecase, Validator: validator}
}

func (h *AttendanceHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	attendances, err := h.Usecase.GetAll(ctx)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	if len(attendances) == 0 {
		SuccessResponse(c, http.StatusOK, "no attendance records found", nil)
		return
	}

	SuccessResponse(c, http.StatusOK, "attendance records retrieved successfully", attendances)
}

func (h *AttendanceHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid attendance ID", nil)
		return
	}

	attendance, err := h.Usecase.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrAttendanceNotFound) {
			ErrorResponse(c, http.StatusNotFound, "attendance record not found", nil)
			return
		}
		ErrorResponse(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	SuccessResponse(c, http.StatusOK, "attendance record retrieved successfully", attendance)
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
		if errors.Is(err, repository.ErrStudentNotFound) {
			ErrorResponse(c, http.StatusNotFound, "student not found", nil)
			return
		}
		if errors.Is(err, repository.ErrAlreadyCheckedIn) {
			ErrorResponse(c, http.StatusConflict, "student already checked in", nil)
			return
		}
		ErrorResponse(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	SuccessResponse(c, http.StatusCreated, "Attendance checked in successfully", &id)
}
