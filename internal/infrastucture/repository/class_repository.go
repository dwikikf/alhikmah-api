package repository

import (
	"context"
	"errors"

	"github.com/dwikikf/alhikmah-api/internal/domain"
	"github.com/dwikikf/alhikmah-api/internal/infrastucture/database"
	repo "github.com/dwikikf/alhikmah-api/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type ClassRepositoryImpl struct {
	db database.DBTX
}

func NewClassRepository(db database.DBTX) *ClassRepositoryImpl {
	return &ClassRepositoryImpl{db: db}
}

func (r *ClassRepositoryImpl) FindAll(ctx context.Context) ([]domain.Class, error) {
	rows, err := r.db.Query(ctx, "SELECT id, code, name, grade FROM classes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classes []domain.Class
	for rows.Next() {
		var class domain.Class
		if err := rows.Scan(&class.ID, &class.Code, &class.Name, &class.Grade); err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return classes, nil
}

func (r *ClassRepositoryImpl) FindByID(ctx context.Context, id int) (*domain.Class, error) {
	row := r.db.QueryRow(ctx, "SELECT id, code, name, grade FROM classes WHERE id = $1", id)

	var class domain.Class
	if err := row.Scan(&class.ID, &class.Code, &class.Name, &class.Grade); err != nil {
		if err == pgx.ErrNoRows {
			return nil, repo.ErrClassNotFound
		}
		return nil, err
	}

	return &class, nil
}

func (r *ClassRepositoryImpl) Create(ctx context.Context, class domain.Class) (int, error) {
	var newID int
	query := `INSERT INTO classes (code, name, grade) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(ctx, query, class.Code, class.Name, class.Grade).Scan(&newID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" { // unique_violation
				return 0, repo.ErrDuplicate
			}
		}
		return 0, err
	}
	return newID, nil
}

func (r *ClassRepositoryImpl) Update(ctx context.Context, class domain.Class) error {
	query := `
		UPDATE classes
		SET code = $1, name = $2, grade = $3
		WHERE id = $4
	`
	cmdTag, err := r.db.Exec(ctx, query, class.Code, class.Name, class.Grade, class.ID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" { // unique_violation
				return repo.ErrDuplicate
			}
		}
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return repo.ErrClassNotFound
	}

	return nil
}

func (r *ClassRepositoryImpl) Delete(ctx context.Context, id int) error {
	cmdTag, err := r.db.Exec(ctx, "DELETE FROM classes WHERE id = $1", id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23503" { // foreign_key_violation
				return repo.ErrForeignKey
			}
		}
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return repo.ErrClassNotFound
	}

	return nil
}
