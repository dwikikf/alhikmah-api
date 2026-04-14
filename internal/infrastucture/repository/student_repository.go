package repository

import (
	"context"

	"github.com/dwikikf/alhikmah-api/internal/domain"
	"github.com/dwikikf/alhikmah-api/internal/infrastucture/database"
	repo "github.com/dwikikf/alhikmah-api/internal/repository"
	"github.com/jackc/pgx/v5"
)

type studentRepository struct {
	db database.DBTX
}

func NewStudentRepository(db database.DBTX) *studentRepository {
	return &studentRepository{db: db}
}

func (r *studentRepository) FindAll(ctx context.Context) ([]domain.Student, error) {
	rows, err := r.db.Query(ctx, "SELECT id, nisn, name FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []domain.Student
	for rows.Next() {
		var student domain.Student
		if err := rows.Scan(&student.ID, &student.NISN, &student.Name); err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (r *studentRepository) FindByID(ctx context.Context, id int) (*domain.Student, error) {
	row := r.db.QueryRow(ctx, "SELECT id, nisn, name FROM students WHERE id = $1", id)

	var student domain.Student
	if err := row.Scan(&student.ID, &student.NISN, &student.Name); err != nil {
		if err == pgx.ErrNoRows {
			return nil, repo.ErrNotFound
		}
		return nil, err
	}

	return &student, nil
}

func (r *studentRepository) Create(ctx context.Context, student domain.Student) (*domain.Student, error) {
	query := `
		INSERT INTO students (nisn, name)
		VALUES ($1, $2)
		RETURNING id
	`
	err := r.db.QueryRow(ctx, query, student.NISN, student.Name).Scan(&student.ID)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (r *studentRepository) Update(ctx context.Context, student domain.Student) error {
	query := `
		UPDATE students
		SET nisn = $1, name = $2
		WHERE id = $3
	`
	cmdTag, err := r.db.Exec(ctx, query, student.NISN, student.Name, student.ID)
	if err != nil {
		// log.Printf("failed to update student with ID %d: %v", student.ID, err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return repo.ErrNotFound
	}

	return nil
}

func (r *studentRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM students WHERE id = $1`

	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return repo.ErrNotFound
	}

	return nil
}
