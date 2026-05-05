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
	"github.com/jackc/pgx/v5/pgconn"
)

type AttendanceRepositoryImpl struct {
	db database.DBTX
}

func NewAttendanceRepository(db database.DBTX) *AttendanceRepositoryImpl {
	return &AttendanceRepositoryImpl{db: db}
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
