package usecase

import (
	"context"
	"time"

	"github.com/dwikikf/alhikmah-api/internal/dto"
	"github.com/dwikikf/alhikmah-api/internal/repository"
)

type AttendanceUsecaseImpl struct {
	StudentRepo repository.StudentRepository
	AttRepo     repository.AttendanceRepository
	uow         repository.UnitOfWork
}

func NewAttendanceUseCaseImpl(
	StudentRepo repository.StudentRepository,
	AttRepo repository.AttendanceRepository,
	uow repository.UnitOfWork,
) *AttendanceUsecaseImpl {
	return &AttendanceUsecaseImpl{
		StudentRepo: StudentRepo,
		AttRepo:     AttRepo,
		uow:         uow,
	}
}

func (uc *AttendanceUsecaseImpl) GetAll(ctx context.Context) ([]dto.AttendanceResponse, error) {
	attendances, err := uc.AttRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return dto.ToAttendanceListResponse(attendances), nil
}

func (uc *AttendanceUsecaseImpl) GetByID(ctx context.Context, id int) (*dto.AttendanceResponse, error) {
	attendance, err := uc.AttRepo.FindByID(ctx, id)
	if err != nil {
		if attendance == nil {
			return nil, repository.ErrAttendanceNotFound
		}
		return nil, err
	}

	return dto.ToAttendanceResponse(*attendance), nil
}

func (uc *AttendanceUsecaseImpl) IsCheckedIn(ctx context.Context, studentID int, date time.Time) (bool, error) {
	var exists bool
	// err := uc.uow.Do(ctx, func(repo repository.Repository) error {
	// 	var err error
	// 	exists, err = repo.Attendance().ExistByStudentAndDate(ctx, studentID, date)
	// 	return err
	// })

	// if err != nil {
	// 	return false, err
	// }

	// return exists, nil

	exists, err := uc.AttRepo.ExistByStudentAndDate(ctx, studentID, date)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (uc *AttendanceUsecaseImpl) CheckIn(ctx context.Context, att dto.CreateAttendanceRequest) (int, error) {
	var id int

	student, err := uc.StudentRepo.FindByID(ctx, att.StudentID)
	if err != nil {
		return 0, err
	}

	if student == nil {
		return 0, repository.ErrStudentNotFound
	}

	attDomain := dto.ToAttendanceDomain(att)

	if attDomain.AttendanceDate == nil {
		now := time.Now()
		attDomain.AttendanceDate = &now
	}

	existing, err := uc.AttRepo.ExistByStudentAndDate(ctx, att.StudentID, *attDomain.AttendanceDate)
	if err != nil {
		return 0, err
	}

	if existing {
		return 0, repository.ErrAlreadyCheckedIn
	}

	err = uc.uow.Do(ctx, func(repo repository.Repository) error {
		var err error
		id, err = repo.Attendance().Create(ctx, attDomain)
		return err
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
