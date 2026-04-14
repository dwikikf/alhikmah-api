package repository

import (
	"context"

	"github.com/dwikikf/alhikmah-api/internal/domain"
)

type StudentRepository interface {
	FindAll(ctx context.Context) ([]domain.Student, error)
	FindByID(ctx context.Context, id int) (*domain.Student, error)
	Create(ctx context.Context, student domain.Student) (*domain.Student, error)
	Update(ctx context.Context, student domain.Student) error
	Delete(ctx context.Context, id int) error
}
