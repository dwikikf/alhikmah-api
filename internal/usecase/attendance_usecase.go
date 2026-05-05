package usecase

import (
	"context"
	"time"

	"github.com/dwikikf/alhikmah-api/internal/dto"
	"github.com/dwikikf/alhikmah-api/internal/repository"
)

type AttendanceUsecaseImpl struct {
	AttRepo repository.AttendanceRepository
	uow     repository.UnitOfWork
}

func NewAttendanceUseCaseImpl(
	AttRepo repository.AttendanceRepository,
	uow repository.UnitOfWork,
) *AttendanceUsecaseImpl {
	return &AttendanceUsecaseImpl{
		AttRepo: AttRepo,
		uow:     uow,
	}
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
