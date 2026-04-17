package usecase

import (
	"context"

	"github.com/dwikikf/alhikmah-api/internal/dto"
)

type StudentUseCase interface {
	GetAllStudents(ctx context.Context) ([]dto.StudentResponse, error)
	GetStudentByID(ctx context.Context, id int) (*dto.StudentResponse, error)
	CreateStudent(ctx context.Context, student dto.CreateStudentRequest) (*dto.StudentResponse, error)
	UpdateStudent(ctx context.Context, id int, student dto.UpdateStudentRequest) (*dto.StudentResponse, error)
	DeleteStudent(ctx context.Context, id int) error
}
