package repository

import "context"

type Repository interface {
	Student() StudentRepository
}

type UnitOfWork interface {
	Do(ctx context.Context, fn func(repo Repository) error) error
}
