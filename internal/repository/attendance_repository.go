package repository

import (
	"context"
	"time"

	"github.com/dwikikf/alhikmah-api/internal/domain"
)

type AttendanceRepository interface {
	FindAll(ctx context.Context) ([]domain.Attendance, error)
	FindByID(ctx context.Context, id int) (*domain.Attendance, error)
	ExistByStudentAndDate(ctx context.Context, studentID int, date time.Time) (bool, error)
	Create(ctx context.Context, attendance domain.Attendance) (int, error)
}
