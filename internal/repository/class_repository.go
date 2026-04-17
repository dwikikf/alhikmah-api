package repository

import (
	"context"

	"github.com/dwikikf/alhikmah-api/internal/domain"
)

type ClassRepository interface {
	FindAll(ctx context.Context) ([]domain.Class, error)
	FindByID(ctx context.Context, id int) (*domain.Class, error)
	Create(ctx context.Context, class domain.Class) (int, error)
	Update(ctx context.Context, class domain.Class) error
	Delete(ctx context.Context, id int) error
}
