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
		SuccessResponse(c, http.StatusOK, "Siswa tidak ditemukan, data kosong", nil)
		return
	}

	SuccessResponse(c, http.StatusOK, "Siswa berhasil diambil", students)
}

func (h *StudentHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "ID siswa tidak valid", nil)
		return
	}

	student, err := h.Usecase.GetStudentByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrStudentNotFound) {
			ErrorResponse(c, http.StatusNotFound, "Siswa tidak ditemukan", nil)
			return
		}

		ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}

	SuccessResponse(c, http.StatusOK, "Siswa berhasil diambil", &student)
}

func (h *StudentHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "ID siswa tidak valid", nil)
		return
	}

	if validationErr := h.validator.Validate(req); validationErr != nil {
		ErrorResponse(c, http.StatusUnprocessableEntity, "Data tidak valid", &validationErr)
		return
	}

	student, err := h.Usecase.CreateStudent(ctx, req)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicate) {
			ErrorResponse(c, http.StatusConflict, "NISN sudah digunakan", nil)
			return
		}

		if errors.Is(err, repository.ErrClassNotFound) {
			ErrorResponse(c, http.StatusNotFound, "Kelas tidak ditemukan", nil)
			return
		}

		ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}

	SuccessResponse(c, http.StatusCreated, "Siswa berhasil dibuat", &student)
}

func (h *StudentHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()

	// validasi ID
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "ID siswa tidak valid", nil)
		return
	}

	var req dto.UpdateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(c, http.StatusBadRequest, "Data permintaan tidak valid", nil)
		return
	}

	if validationErr := h.validator.Validate(req); validationErr != nil {
		ErrorResponse(c, http.StatusUnprocessableEntity, "Data tidak valid", &validationErr)
		return
	}

	err = h.Usecase.UpdateStudent(ctx, id, req)
	if err != nil {
		if errors.Is(err, repository.ErrClassNotFound) {
			ErrorResponse(c, http.StatusNotFound, "Kelas tidak ditemukan", nil)
			return
		}

		if errors.Is(err, repository.ErrStudentNotFound) {
			ErrorResponse(c, http.StatusNotFound, "Siswa tidak ditemukan", nil)
			return
		}

		if errors.Is(err, repository.ErrDuplicate) {
			ErrorResponse(c, http.StatusConflict, "NISN sudah digunakan", nil)
			return
		}

		if errors.Is(err, repository.ErrForeignKey) {
			ErrorResponse(c, http.StatusBadRequest, "Kelas tidak ditemukan", nil)
			return
		}

		ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}

	SuccessResponse(c, http.StatusOK, "Siswa berhasil diperbarui", nil)
}

func (h *StudentHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, "ID siswa tidak valid", nil)
		return
	}

	err = h.Usecase.DeleteStudent(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrStudentNotFound) {
			ErrorResponse(c, http.StatusNotFound, "Siswa tidak ditemukan", nil)
			return
		}

		if errors.Is(err, repository.ErrForeignKey) {
			ErrorResponse(c, http.StatusConflict, "Gagal menghapus data. Data ini masih digunakan oleh data atau fitur lain.", nil)
			return
		}

		ErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}

	SuccessResponse(c, http.StatusOK, "Siswa berhasil dihapus", nil)
}
