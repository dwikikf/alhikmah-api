package usecase

import (
	"context"

	"github.com/dwikikf/alhikmah-api/internal/domain"
	"github.com/dwikikf/alhikmah-api/internal/repository"
)

type StudentUseCase struct {
	StudentRepo repository.StudentRepository // untuk GET
	Uow         repository.UnitOfWork        // untuk WRITE
}

func NewStudentUseCase(
	studentRepo repository.StudentRepository,
	uow repository.UnitOfWork,
) *StudentUseCase {
	return &StudentUseCase{
		StudentRepo: studentRepo,
		Uow:         uow,
	}
}

func (u *StudentUseCase) GetAllStudents(ctx context.Context) ([]domain.Student, error) {
	return u.StudentRepo.FindAll(ctx)
}

func (u *StudentUseCase) GetStudentByID(ctx context.Context, id int) (*domain.Student, error) {
	return u.StudentRepo.FindByID(ctx, id)
}

func (u *StudentUseCase) CreateStudent(ctx context.Context, student domain.Student) (*domain.Student, error) {
	var createdStudent *domain.Student

	err := u.Uow.Do(ctx, func(repo repository.Repository) error {
		var err error
		createdStudent, err = repo.Student().Create(ctx, student)
		return err
	})

	return createdStudent, err
}

func (u *StudentUseCase) UpdateStudent(ctx context.Context, student domain.Student) error {
	return u.Uow.Do(ctx, func(repo repository.Repository) error {

		// cek apakah student dengan ID ini ada?
		// find by ID ini dilakukan untuk kedepanya jika memang ada bussiness logic yang perlu dilakukan sebelum update
		existing, err := repo.Student().FindByID(ctx, student.ID)
		if err != nil {
			return err
		}

		if existing == nil {
			return repository.ErrNotFound
		}

		return repo.Student().Update(ctx, student)
	})
}

func (u *StudentUseCase) DeleteStudent(ctx context.Context, id int) error {
	return u.Uow.Do(ctx, func(repo repository.Repository) error {
		return repo.Student().Delete(ctx, id)
	})
}
