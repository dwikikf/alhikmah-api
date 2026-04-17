package repository

import (
	repo "github.com/dwikikf/alhikmah-api/internal/repository"
	"github.com/jackc/pgx/v5"
)

type repository struct {
	tx pgx.Tx
}

func NewRepositoryWithTx(tx pgx.Tx) repo.Repository {
	return &repository{tx: tx}
}

func (r *repository) Student() repo.StudentRepository {
	return NewStudentRepository(r.tx)
}

func (r *repository) Class() repo.ClassRepository {
	return NewClassRepository(r.tx)
}
