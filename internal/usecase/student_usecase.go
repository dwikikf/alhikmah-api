package usecase

import (
	"context"

	"github.com/dwikikf/alhikmah-api/internal/domain"
	"github.com/dwikikf/alhikmah-api/internal/dto"
	"github.com/dwikikf/alhikmah-api/internal/repository"
)

type StudentUsecase struct {
	StudentRepo repository.StudentRepository // untuk GET
	Uow         repository.UnitOfWork        // untuk WRITE
}

func NewStudentUseCase(
	studentRepo repository.StudentRepository,
	uow repository.UnitOfWork,
) *StudentUsecase {
	return &StudentUsecase{
		StudentRepo: studentRepo,
		Uow:         uow,
	}
}

func (u *StudentUsecase) GetAllStudents(ctx context.Context) ([]dto.StudentResponse, error) {
	students, err := u.StudentRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return dto.ToStudentListResponse(students), nil

}

func (u *StudentUsecase) GetStudentByID(ctx context.Context, id int) (*dto.StudentResponse, error) {
	student, err := u.StudentRepo.FindByID(ctx, id)
	if err != nil {
		if student == nil {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return &dto.StudentResponse{
		ID:   student.ID,
		NISN: student.NISN,
		Name: student.Name,
	}, nil
}

func (u *StudentUsecase) CreateStudent(ctx context.Context, student dto.CreateStudentRequest) (*dto.StudentResponse, error) {
	var createdStudent *domain.Student
	studentDomain := dto.ToCreateStudentDomain(student)

	err := u.Uow.Do(ctx, func(repo repository.Repository) error {
		var err error
		createdStudent, err = repo.Student().Create(ctx, studentDomain)
		return err
	})

	if err != nil {
		return nil, err
	}

	return &dto.StudentResponse{
		ID:   createdStudent.ID,
		NISN: createdStudent.NISN,
		Name: createdStudent.Name,
	}, err
}

func (u *StudentUsecase) UpdateStudent(ctx context.Context, id int, student dto.UpdateStudentRequest) error {
	studentDomain := dto.ToUpdateStudentDomain(id, student)

	return u.Uow.Do(ctx, func(repo repository.Repository) error {

		// cek apakah student dengan ID ini ada?
		// find by ID ini dilakukan untuk kedepanya jika memang ada bussiness logic yang perlu dilakukan sebelum update
		existing, err := repo.Student().FindByID(ctx, id)

		if err != nil {
			return err
		}

		if existing == nil {
			return repository.ErrNotFound
		}

		return repo.Student().Update(ctx, studentDomain)
	})
}

func (u *StudentUsecase) DeleteStudent(ctx context.Context, id int) error {
	return u.Uow.Do(ctx, func(repo repository.Repository) error {
		return repo.Student().Delete(ctx, id)
	})
}
