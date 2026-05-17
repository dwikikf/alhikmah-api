package dto

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

type StudentResponse struct {
	ID    int           `json:"id"`
	NISN  string        `json:"nisn"`
	Name  string        `json:"name"`
	Class ClassResponse `json:"class"`
}

type ClassResponse struct {
	ID        int     `json:"id"`
	Code      string  `json:"code"`
	Name      string  `json:"name"`
	Grade     int     `json:"grade"`
	StartTime *string `json:"start_time"`
}

type AttendanceResponse struct {
	ID int `json:"id"`
	// StudentID      int             `json:"student_id"`
	Student        StudentResponse `json:"student"`
	AttendanceDate string          `json:"attendance_date"`
	CheckIn        string          `json:"check_in"`
	Status         string          `json:"status"`
	Method         string          `json:"method"`
	Note           *string         `json:"note,omitempty"`
	IsLate         *bool           `json:"is_late,omitempty"`
}
