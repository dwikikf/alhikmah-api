package usecase

import (
	"context"

	"github.com/dwikikf/alhikmah-api/internal/domain"
	"github.com/dwikikf/alhikmah-api/internal/dto"
	"github.com/dwikikf/alhikmah-api/internal/repository"
)

type StudentUsecaseImpl struct {
	ClassRepo   repository.ClassRepository   // untuk cek class saat create/update student
	StudentRepo repository.StudentRepository // untuk GET
	Uow         repository.UnitOfWork        // untuk WRITE
}

func NewStudentUseCaseImpl(
	classRepo repository.ClassRepository,
	studentRepo repository.StudentRepository,
	uow repository.UnitOfWork,
) *StudentUsecaseImpl {
	return &StudentUsecaseImpl{
		ClassRepo:   classRepo,
		StudentRepo: studentRepo,
		Uow:         uow,
	}
}

func (u *StudentUsecaseImpl) GetAllStudents(ctx context.Context) ([]dto.StudentResponse, error) {
	students, err := u.StudentRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return dto.ToStudentListResponse(students), nil

}

func (u *StudentUsecaseImpl) GetStudentByID(ctx context.Context, id int) (*dto.StudentResponse, error) {
	student, err := u.StudentRepo.FindByID(ctx, id)
	if err != nil {
		if student == nil {
			return nil, repository.ErrStudentNotFound
		}
		return nil, err
	}

	return dto.ToStudentResponse(*student), nil
}

func (u *StudentUsecaseImpl) CreateStudent(ctx context.Context, student dto.CreateStudentRequest) (*dto.StudentResponse, error) {
	class, err := u.ClassRepo.FindByID(ctx, student.ClassID)
	if err != nil {
		if class == nil {
			return nil, repository.ErrClassNotFound
		}
		return nil, err
	}

	var createdStudent *domain.Student
	studentDomain := dto.ToCreateStudentDomain(student)

	err = u.Uow.Do(ctx, func(repo repository.Repository) error {
		var err error
		createdStudent, err = repo.Student().Create(ctx, studentDomain)
		return err
	})

	if err != nil {
		return nil, err
	}

	// return dto.ToStudentResponse(*createdStudent), nil
	return &dto.StudentResponse{
		ID:   createdStudent.ID,
		NISN: createdStudent.NISN,
		Name: createdStudent.Name,
		Class: dto.ClassResponse{
			ID:        class.ID,
			Code:      class.Code,
			Name:      class.Name,
			Grade:     class.Grade,
			StartTime: class.StartTime,
		},
	}, err
}

func (u *StudentUsecaseImpl) UpdateStudent(ctx context.Context, id int, student dto.UpdateStudentRequest) error {
	class, err := u.ClassRepo.FindByID(ctx, student.ClassID)
	if err != nil {
		if class == nil {
			return repository.ErrClassNotFound
		}
		return err
	}

	studentDomain := dto.ToUpdateStudentDomain(id, student)

	return u.Uow.Do(ctx, func(repo repository.Repository) error {

		// cek apakah student dengan ID ini ada?
		// find by ID ini dilakukan untuk kedepanya jika memang ada bussiness logic yang perlu dilakukan sebelum update
		existing, err := repo.Student().FindByID(ctx, id)

		if err != nil {
			return err
		}

		if existing == nil {
			return repository.ErrStudentNotFound
		}

		return repo.Student().Update(ctx, studentDomain)
	})
}

func (u *StudentUsecaseImpl) DeleteStudent(ctx context.Context, id int) error {
	return u.Uow.Do(ctx, func(repo repository.Repository) error {
		return repo.Student().Delete(ctx, id)
	})
}
