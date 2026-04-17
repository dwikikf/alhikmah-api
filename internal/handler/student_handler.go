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
		c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  "error",
			Message: "internal server error",
		})
		return
	}

	// optional: handle empty data
	if len(students) == 0 {
		c.JSON(http.StatusOK, dto.Response{
			Status:  "success",
			Message: "no students found",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  "success",
		Message: "students retrieved successfully",
		Data:    students,
	})
}

func (h *StudentHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Status:  "error",
			Message: "invalid student ID",
		})
		return
	}

	student, err := h.Usecase.GetStudentByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  "error",
			Message: "internal server error",
		})
		return
	}

	if student == nil {
		c.JSON(http.StatusNotFound, dto.Response{
			Status:  "error",
			Message: "student not found",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  "success",
		Message: "student retrieved successfully",
		Data:    student,
	})
}

func (h *StudentHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Status:  "error",
			Message: "invalid request body",
		})
		return
	}

	newStudent := dto.CreateStudentRequest{
		NISN: req.NISN,
		Name: req.Name,
	}

	student, err := h.Usecase.CreateStudent(ctx, newStudent)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicate) {
			c.JSON(http.StatusConflict, dto.Response{
				Status:  "error",
				Message: "NISN already exists",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  "error",
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Status:  "success",
		Message: "student created successfully",
		Data:    student,
	})
}

func (h *StudentHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()

	// validasi ID
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Status:  "error",
			Message: "invalid student ID",
		})
		return
	}

	var req dto.UpdateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Status:  "error",
			Message: "invalid request body",
		})
		return
	}

	updatedStudent := dto.UpdateStudentRequest{
		NISN: req.NISN,
		Name: req.Name,
	}

	err = h.Usecase.UpdateStudent(ctx, id, updatedStudent)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			c.JSON(http.StatusNotFound, dto.Response{
				Status:  "error",
				Message: "student not found",
			})
			return
		}

		if errors.Is(err, repository.ErrDuplicate) {
			c.JSON(http.StatusConflict, dto.Response{
				Status:  "error",
				Message: "NISN already exists",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  "error",
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  "success",
		Message: "student updated successfully",
	})
}

func (h *StudentHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Status:  "error",
			Message: "invalid student ID",
		})
		return
	}

	err = h.Usecase.DeleteStudent(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			c.JSON(http.StatusNotFound, dto.Response{
				Status:  "error",
				Message: "student not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, dto.Response{
			Status:  "error",
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  "success",
		Message: "student deleted successfully",
	})
}
