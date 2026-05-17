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

type StudentHandler struct {
	Usecase   *usecase.StudentUsecaseImpl
	validator *validator.CustomValidator
}

func NewStudentHandler(usecase *usecase.StudentUsecaseImpl, validator *validator.CustomValidator) *StudentHandler {
	return &StudentHandler{
		Usecase:   usecase,
		validator: validator,
	}
}

func (h *StudentHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	students, err := h.Usecase.GetAllStudents(ctx)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}

	if len(students) == 0 {
		SuccessResponse(c, http.StatusOK, "No students found", nil)
		return
	}

	SuccessResponse(c, http.StatusOK, "Students retrieved successfully", students)
}

func (h *StudentHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Invalid student ID", nil)
		return
	}

	student, err := h.Usecase.GetStudentByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrStudentNotFound) {
			ErrorResponse(c, http.StatusNotFound, "Student not found", nil)
			return
		}

		ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}

	SuccessResponse(c, http.StatusOK, "Student retrieved successfully", &student)
}

func (h *StudentHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Invalid student ID", nil)
		return
	}

	if validationErr := h.validator.Validate(req); validationErr != nil {
		ErrorResponse(c, http.StatusUnprocessableEntity, "Invalid student data", &validationErr)
		return
	}

	student, err := h.Usecase.CreateStudent(ctx, req)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicate) {
			ErrorResponse(c, http.StatusConflict, "NISN already exists", nil)
			return
		}

		if errors.Is(err, repository.ErrClassNotFound) {
			ErrorResponse(c, http.StatusNotFound, "Class not found", nil)
			return
		}

		ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}

	SuccessResponse(c, http.StatusCreated, "Student created successfully", &student)
}

func (h *StudentHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()

	// validasi ID
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Invalid student ID", nil)
		return
	}

	var req dto.UpdateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Invalid request data", nil)
		return
	}

	if validationErr := h.validator.Validate(req); validationErr != nil {
		ErrorResponse(c, http.StatusUnprocessableEntity, "Invalid student data", &validationErr)
		return
	}

	err = h.Usecase.UpdateStudent(ctx, id, req)
	if err != nil {
		if errors.Is(err, repository.ErrClassNotFound) {
			ErrorResponse(c, http.StatusNotFound, "Class not found", nil)
			return
		}

		if errors.Is(err, repository.ErrStudentNotFound) {
			ErrorResponse(c, http.StatusNotFound, "Student not found", nil)
			return
		}

		if errors.Is(err, repository.ErrDuplicate) {
			ErrorResponse(c, http.StatusConflict, "NISN already exists", nil)
			return
		}

		if errors.Is(err, repository.ErrForeignKey) {
			ErrorResponse(c, http.StatusBadRequest, "Class not found", nil)
			return
		}

		ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}

	SuccessResponse(c, http.StatusOK, "Student updated successfully", nil)
}

func (h *StudentHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Invalid student ID", nil)
		return
	}

	err = h.Usecase.DeleteStudent(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrStudentNotFound) {
			ErrorResponse(c, http.StatusNotFound, "Student not found", nil)
			return
		}

		if errors.Is(err, repository.ErrForeignKey) {
			ErrorResponse(c, http.StatusConflict, "Failed to delete data. This data is still being used by other data or features.", nil)
			return
		}

		ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}

	SuccessResponse(c, http.StatusOK, "Student deleted successfully", nil)
}
