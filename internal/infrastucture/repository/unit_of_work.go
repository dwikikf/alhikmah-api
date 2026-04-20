package repository

import (
	"context"

	repo "github.com/dwikikf/alhikmah-api/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UnitOfWorkImpl struct {
	db *pgxpool.Pool
}

func NewUnitOfWork(db *pgxpool.Pool) *UnitOfWorkImpl {
	return &UnitOfWorkImpl{db: db}
}

func (u *UnitOfWorkImpl) Do(ctx context.Context, fn func(repo.Repository) error) error {
	tx, err := u.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// 🔥 buat repo berbasis TX
	r := NewRepositoryWithTx(tx)

	// 🔥 kirim interface, bukan struct
	if err := fn(r); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
