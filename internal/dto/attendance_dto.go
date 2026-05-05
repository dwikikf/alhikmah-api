package dto

import "github.com/dwikikf/alhikmah-api/internal/domain"

func ToAttendanceResponse(attendance domain.Attendance) *AttendanceResponse {
	return &AttendanceResponse{
		ID:     attendance.ID,
		Status: string(attendance.Status),
		Method: string(attendance.Method),
	}
}

func ToAttendanceDomain(req CreateAttendanceRequest) domain.Attendance {
	return domain.Attendance{
		StudentID:      req.StudentID,
		AttendanceDate: req.AttendanceDate,
		Status:         domain.AttendanceStatus(req.Status),
		Method:         domain.AttendanceStatus(req.Method),
		Note:           &req.Note,
	}
}
