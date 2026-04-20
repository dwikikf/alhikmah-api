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
	Usecase   *usecase.StudentUsecase
	validator *validator.CustomValidator
}

func NewStudentHandler(usecase *usecase.StudentUsecase, validator *validator.CustomValidator) *StudentHandler {
	return &StudentHandler{
		Usecase:   usecase,
		validator: validator,
	}
}

func (h *StudentHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	students, err := h.Usecase.GetAllStudents(ctx)
	if err != nil {
		ErrorResponse[any](c, http.StatusInternalServerError, "Kesalahan sistem internal", nil)
		return
	}

	if len(students) == 0 {
		SuccessResponse[any](c, http.StatusOK, "Siswa tidak ditemukan, data kosong", nil)
		return
	}

	SuccessResponse(c, http.StatusOK, "Siswa berhasil diambil", &students)
}

func (h *StudentHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ErrorResponse[any](c, http.StatusBadRequest, "ID siswa tidak valid", nil)
		return
	}

	student, err := h.Usecase.GetStudentByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrStudentNotFound) {
			ErrorResponse[any](c, http.StatusNotFound, "Siswa tidak ditemukan", nil)
			return
		}

		ErrorResponse[any](c, http.StatusInternalServerError, "Kesalahan sistem internal", nil)
		return
	}

	SuccessResponse(c, http.StatusOK, "Siswa berhasil diambil", student)
}

func (h *StudentHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse[any](c, http.StatusBadRequest, "ID siswa tidak valid", nil)
		return
	}

	if validationErr := h.validator.Validate(req); validationErr != nil {
		ErrorResponse(c, http.StatusUnprocessableEntity, "Data tidak valid", &validationErr)
		return
	}

	student, err := h.Usecase.CreateStudent(ctx, req)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicate) {
			ErrorResponse[any](c, http.StatusConflict, "NISN sudah digunakan", nil)
			return
		}

		if errors.Is(err, repository.ErrStudentNotFound) {
			ErrorResponse[any](c, http.StatusNotFound, "Kelas tidak ditemukan", nil)
			return
		}

		ErrorResponse[any](c, http.StatusInternalServerError, "Kesalahan sistem internal", nil)
		return
	}

	SuccessResponse(c, http.StatusCreated, "Siswa berhasil dibuat", student)
}

func (h *StudentHandler) Update(c *gin.Context) {
	ctx := c.Request.Context()

	// validasi ID
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ErrorResponse[any](c, http.StatusBadRequest, "ID siswa tidak valid", nil)
		return
	}

	var req dto.UpdateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse[any](c, http.StatusBadRequest, "Data permintaan tidak valid", nil)
		return
	}

	if validationErr := h.validator.Validate(req); validationErr != nil {
		ErrorResponse(c, http.StatusUnprocessableEntity, "Data tidak valid", &validationErr)
		return
	}

	err = h.Usecase.UpdateStudent(ctx, id, req)
	if err != nil {
		if errors.Is(err, repository.ErrClassNotFound) {
			ErrorResponse[any](c, http.StatusNotFound, "Kelas tidak ditemukan", nil)
			return
		}

		if errors.Is(err, repository.ErrStudentNotFound) {
			ErrorResponse[any](c, http.StatusNotFound, "Siswa tidak ditemukan", nil)
			return
		}

		if errors.Is(err, repository.ErrDuplicate) {
			ErrorResponse[any](c, http.StatusConflict, "NISN sudah digunakan", nil)
			return
		}

		if errors.Is(err, repository.ErrForeignKey) {
			ErrorResponse[any](c, http.StatusBadRequest, "Kelas tidak ditemukan", nil)
			return
		}

		ErrorResponse[any](c, http.StatusInternalServerError, "Kesalahan sistem internal", nil)
		return
	}

	SuccessResponse[any](c, http.StatusOK, "Siswa berhasil diperbarui", nil)
}

func (h *StudentHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ErrorResponse[any](c, http.StatusBadRequest, "ID siswa tidak valid", nil)
		return
	}

	err = h.Usecase.DeleteStudent(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrStudentNotFound) {
			ErrorResponse[any](c, http.StatusNotFound, "Siswa tidak ditemukan", nil)
			return
		}

		ErrorResponse[any](c, http.StatusInternalServerError, "Kesalahan sistem internal", nil)
		return
	}

	SuccessResponse[any](c, http.StatusOK, "Siswa berhasil dihapus", nil)
}
