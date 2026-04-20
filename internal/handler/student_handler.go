package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/dwikikf/alhikmah-api/internal/dto"
	"github.com/dwikikf/alhikmah-api/internal/repository"
	"github.com/dwikikf/alhikmah-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	Usecase *usecase.StudentUsecase
}

func NewStudentHandler(usecase *usecase.StudentUsecase) *StudentHandler {
	return &StudentHandler{
		Usecase: usecase,
	}
}

func (h *StudentHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	students, err := h.Usecase.GetAllStudents(ctx)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	// optional: handle empty data
	if len(students) == 0 {
		SuccessResponse[dto.StudentResponse](c, http.StatusOK, "no students found", nil)
		return
	}

	SuccessResponse(c, http.StatusOK, "students retrieved successfully", &students)
}

func (h *StudentHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid student ID")
		return
	}

	student, err := h.Usecase.GetStudentByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			ErrorResponse(c, http.StatusNotFound, "student not found")
			return
		}

		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	SuccessResponse(c, http.StatusOK, "student retrieved successfully", student)
}

func (h *StudentHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	newStudent := dto.CreateStudentRequest{
		NISN:    req.NISN,
		Name:    req.Name,
		ClassID: req.ClassID,
	}

	student, err := h.Usecase.CreateStudent(ctx, newStudent)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicate) {
			ErrorResponse(c, http.StatusConflict, "NISN already exists")
			return
		}
		log.Printf("error creating student: %v", err)
		if errors.Is(err, repository.ErrNotFound) {
			ErrorResponse(c, http.StatusNotFound, "class not found")
			return
		}

		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	SuccessResponse(c, http.StatusCreated, "student created successfully", student)
}

func (h *StudentHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()

	// validasi ID
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid student ID")
		return
	}

	var req dto.UpdateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	updatedStudent := dto.UpdateStudentRequest{
		NISN: req.NISN,
		Name: req.Name,
	}

	err = h.Usecase.UpdateStudent(ctx, id, updatedStudent)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			ErrorResponse(c, http.StatusNotFound, "student not found")
			return
		}

		if errors.Is(err, repository.ErrDuplicate) {
			ErrorResponse(c, http.StatusConflict, "NISN already exists")
			return
		}

		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	SuccessResponse[dto.StudentResponse](c, http.StatusOK, "student updated successfully", nil)
}

func (h *StudentHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "invalid student ID")
		return
	}

	err = h.Usecase.DeleteStudent(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			ErrorResponse(c, http.StatusNotFound, "student not found")
			return
		}

		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	SuccessResponse[dto.StudentResponse](c, http.StatusOK, "student deleted successfully", nil)
}
