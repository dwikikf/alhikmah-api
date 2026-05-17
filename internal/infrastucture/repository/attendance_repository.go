package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dwikikf/alhikmah-api/internal/domain"
	"github.com/dwikikf/alhikmah-api/internal/infrastucture/database"
	repo "github.com/dwikikf/alhikmah-api/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type AttendanceRepositoryImpl struct {
	db database.DBTX
}

func NewAttendanceRepository(db database.DBTX) *AttendanceRepositoryImpl {
	return &AttendanceRepositoryImpl{db: db}
}

func (r *AttendanceRepositoryImpl) FindAll(ctx context.Context) ([]domain.Attendance, error) {
	query := `
		SELECT a.id, s.id, s.nisn, s.name, a.attendance_date, a.check_in, a.status, a.method, a.note, a.is_late, c.id, c.code, c.name, c.grade, c.start_time
			FROM attendances a
			INNER JOIN students s ON a.student_id = s.id
			INNER JOIN classes c ON s.class_id = c.id
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attendances []domain.Attendance
	for rows.Next() {
		var att domain.Attendance
		if err := rows.Scan(
			&att.ID,
			&att.Student.ID,
			&att.Student.NISN,
			&att.Student.Name,
			&att.AttendanceDate,
			&att.CheckIn,
			&att.Status,
			&att.Method,
			&att.Note,
			&att.IsLate,
			&att.Student.Class.ID,
			&att.Student.Class.Code,
			&att.Student.Class.Name,
			&att.Student.Class.Grade,
			&att.Student.Class.StartTime,
		); err != nil {
			return nil, err
		}
		attendances = append(attendances, att)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return attendances, nil
}

func (r *AttendanceRepositoryImpl) FindByID(ctx context.Context, id int) (*domain.Attendance, error) {
	query := `
		SELECT a.id, s.id, s.nisn, s.name, a.attendance_date, a.check_in, a.status, a.method, a.note, a.is_late, c.id, c.code, c.name, c.grade, c.start_time
			FROM attendances a
			INNER JOIN students s ON a.student_id = s.id
			INNER JOIN classes c ON s.class_id = c.id
		WHERE a.id = $1
	`
	row := r.db.QueryRow(ctx, query, id)

	var att domain.Attendance
	if err := row.Scan(
		&att.ID,
		&att.Student.ID,
		&att.Student.NISN,
		&att.Student.Name,
		&att.AttendanceDate,
		&att.CheckIn,
		&att.Status,
		&att.Method,
		&att.Note,
		&att.IsLate,
		&att.Student.Class.ID,
		&att.Student.Class.Code,
		&att.Student.Class.Name,
		&att.Student.Class.Grade,
		&att.Student.Class.StartTime,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repo.ErrAttendanceNotFound
		}
		return nil, err
	}

	return &att, nil
}

func (r *AttendanceRepositoryImpl) ExistByStudentAndDate(
	ctx context.Context,
	studentID int,
	date time.Time,
) (bool, error) {

	query := `
		SELECT EXISTS (
			SELECT 1
			FROM attendances
			WHERE student_id = $1
			AND attendance_date = DATE($2)
		)
	`

	var exists bool
	err := r.db.QueryRow(ctx, query, studentID, date.Format("2006-01-02")).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *AttendanceRepositoryImpl) Create(
	ctx context.Context,
	att domain.Attendance,
) (int, error) {

	var newID int

	columns := []string{"student_id"}
	values := []any{att.StudentID}
	placeholders := []string{"$1"}

	// dynamic field (optional)
	if att.Status != "" {
		columns = append(columns, "status")
		values = append(values, att.Status)
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(values)))
	}

	if att.Method != "" {
		columns = append(columns, "method")
		values = append(values, att.Method)
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(values)))
	}

	if att.Note != nil {
		columns = append(columns, "note")
		values = append(values, att.Note)
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(values)))
	}

	if att.IsLate != nil {
		columns = append(columns, "is_late")
		values = append(values, att.IsLate)
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(values)))
	}

	query := fmt.Sprintf(`
		INSERT INTO attendances (%s)
		VALUES (%s)
		RETURNING id
	`, strings.Join(columns, ", "), strings.Join(placeholders, ", "))

	err := r.db.QueryRow(ctx, query, values...).Scan(&newID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {

			// unique constraint (student_id + attendance_date)
			if pgErr.Code == "23505" {
				return 0, repo.ErrAlreadyCheckedIn
			}
		}
		return 0, err
	}

	return newID, nil
}
