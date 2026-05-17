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

type StudentRepositoryImpl struct {
	db database.DBTX
}

func NewStudentRepository(db database.DBTX) *StudentRepositoryImpl {
	return &StudentRepositoryImpl{db: db}
}

func (r *StudentRepositoryImpl) FindAll(ctx context.Context) ([]domain.Student, error) {
	query := "SELECT s.id, s.nisn, s.name, c.id as class_id, c.code as class_code, c.name as class_name,  c.grade as class_grade, c.start_time as start_time FROM students s INNER JOIN classes c ON s.class_id = c.id"
	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []domain.Student
	for rows.Next() {
		var student domain.Student
		student.Class = domain.Class{}

		if err := rows.Scan(&student.ID, &student.NISN, &student.Name,
			&student.Class.ID, &student.Class.Code, &student.Class.Name, &student.Class.Grade, &student.Class.StartTime); err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (r *StudentRepositoryImpl) FindByID(ctx context.Context, id int) (*domain.Student, error) {
	query := "SELECT s.id, s.nisn, s.name, c.id as class_id, c.code as class_code, c.name as class_name, c.grade as class_grade, c.start_time as start_time FROM students s INNER JOIN classes c ON s.class_id = c.id WHERE s.id = $1"
	row := r.db.QueryRow(ctx, query, id)

	var student domain.Student
	student.Class = domain.Class{}

	if err := row.Scan(
		&student.ID,
		&student.NISN,
		&student.Name,
		&student.Class.ID,
		&student.Class.Code,
		&student.Class.Name,
		&student.Class.Grade,
		&student.Class.StartTime,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, repo.ErrStudentNotFound
		}
		return nil, err
	}

	return &student, nil
}

func (r *StudentRepositoryImpl) Create(ctx context.Context, student domain.Student) (*domain.Student, error) {
	query := `
		INSERT INTO students (nisn, name, class_id)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err := r.db.QueryRow(ctx, query, student.NISN, student.Name, student.ClassID).Scan(&student.ID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" { // unique_violation
				return nil, repo.ErrDuplicate
			}
		}
		return nil, err
	}

	return &student, nil
}

func (r *StudentRepositoryImpl) Update(ctx context.Context, student domain.Student) error {
	query := `
		UPDATE students
		SET nisn = $1, name = $2, class_id = $3
		WHERE id = $4
	`
	cmdTag, err := r.db.Exec(ctx, query, student.NISN, student.Name, student.ClassID, student.ID)
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
		return repo.ErrStudentNotFound
	}

	return nil
}

func (r *StudentRepositoryImpl) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM students WHERE id = $1`

	cmdTag, err := r.db.Exec(ctx, query, id)
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
		return repo.ErrStudentNotFound
	}

	return nil
}
