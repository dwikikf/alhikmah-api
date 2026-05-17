package dto

import "github.com/dwikikf/alhikmah-api/internal/domain"

func ToAttendanceListResponse(attendances []domain.Attendance) []AttendanceResponse {
	responses := make([]AttendanceResponse, len(attendances))
	for i, attendance := range attendances {
		responses[i] = *ToAttendanceResponse(attendance)
	}
	return responses
}

func ToAttendanceResponse(attendance domain.Attendance) *AttendanceResponse {
	return &AttendanceResponse{
		ID: attendance.ID,
		// StudentID: attendance.StudentID,
		Student: StudentResponse{
			ID:   attendance.Student.ID,
			NISN: attendance.Student.NISN,
			Name: attendance.Student.Name,
			Class: ClassResponse{
				ID:        attendance.Student.Class.ID,
				Code:      attendance.Student.Class.Code,
				Name:      attendance.Student.Class.Name,
				Grade:     attendance.Student.Class.Grade,
				StartTime: attendance.Student.Class.StartTime,
			},
		},
		AttendanceDate: attendance.AttendanceDate.Format("2006-01-02"),
		CheckIn:        attendance.CheckIn.Format("15:04:05"),
		Status:         string(attendance.Status),
		Method:         string(attendance.Method),
		Note:           attendance.Note,
		IsLate:         attendance.IsLate,
	}
}

func ToAttendanceDomain(req CreateAttendanceRequest) domain.Attendance {
	return domain.Attendance{
		StudentID:      req.StudentID,
		AttendanceDate: req.AttendanceDate,
		Status:         domain.AttendanceStatus(req.Status),
		Method:         domain.AttendanceMethod(req.Method),
		Note:           &req.Note,
	}
}
