package repository

import (
	"context"
	"time"

	"github.com/dwikikf/alhikmah-api/internal/domain"
)

type AttendanceRepository interface {
	ExistByStudentAndDate(ctx context.Context, studentID int, date time.Time) (bool, error)
	Create(ctx context.Context, attendance domain.Attendance) (int, error)
}
