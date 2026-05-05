package domain

import (
	"time"
)

type AttendanceStatus string

const (
	StatusHadir AttendanceStatus = "hadir"
	StatusIzin  AttendanceStatus = "izin"
	StatusSakit AttendanceStatus = "sakit"
	StatusAlpha AttendanceStatus = "alpha"
)

type AttendanceMethod string

const (
	MethodQR     AttendanceMethod = "qr"
	MethodManual AttendanceMethod = "manual"
	MethodAdmin  AttendanceMethod = "admin"
)

type Attendance struct {
	ID             int
	StudentID      int
	Student        *Student
	AttendanceDate *time.Time
	CheckIn        time.Time
	CheckOut       *time.Time
	Status         AttendanceStatus
	Method         AttendanceStatus
	Note           *string
	IsLate         *bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
