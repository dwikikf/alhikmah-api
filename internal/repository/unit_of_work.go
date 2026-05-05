package repository

import "context"

type Repository interface {
	Class() ClassRepository
	Student() StudentRepository
	Attendance() AttendanceRepository
}

type UnitOfWork interface {
	Do(ctx context.Context, fn func(repo Repository) error) error
}
