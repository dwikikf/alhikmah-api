package usecase

import (
	"context"

	"github.com/dwikikf/alhikmah-api/internal/dto"
)

type StudentUseCase interface {
	GetAllStudents(ctx context.Context) ([]dto.StudentResponse, error)
	GetStudentByID(ctx context.Context, id int) (*dto.StudentResponse, error)
	CreateStudent(ctx context.Context, student dto.CreateStudentRequest) (int, error)
	UpdateStudent(ctx context.Context, id int, student dto.UpdateStudentRequest) error
	DeleteStudent(ctx context.Context, id int) error
}

type ClassUseCase interface {
	GetAllClasses(ctx context.Context) ([]dto.ClassResponse, error)
	GetClassByID(ctx context.Context, id int) (*dto.ClassResponse, error)
	CreateClass(ctx context.Context, class dto.CreateClassRequest) (int, error)
	UpdateClass(ctx context.Context, id int, class dto.UpdateClassRequest) error
	DeleteClass(ctx context.Context, id int) error
}
