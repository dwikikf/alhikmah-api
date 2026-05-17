package dto

import "github.com/dwikikf/alhikmah-api/internal/domain"

func ToStudentListResponse(students []domain.Student) []StudentResponse {
	var res []StudentResponse
	for _, s := range students {
		res = append(res, StudentResponse{
			ID:   s.ID,
			NISN: s.NISN,
			Name: s.Name,
			// ClassID: s.ClassID,
			Class: ClassResponse{
				ID:        s.Class.ID,
				Code:      s.Class.Code,
				Name:      s.Class.Name,
				Grade:     s.Class.Grade,
				StartTime: s.Class.StartTime,
			},
		})
	}
	return res
}

func ToStudentResponse(students domain.Student) *StudentResponse {
	return &StudentResponse{
		ID:   students.ID,
		NISN: students.NISN,
		Name: students.Name,
		Class: ClassResponse{
			ID:        students.Class.ID,
			Code:      students.Class.Code,
			Name:      students.Class.Name,
			Grade:     students.Class.Grade,
			StartTime: students.Class.StartTime,
		},
	}
}

func ToCreateStudentDomain(req CreateStudentRequest) domain.Student {
	return domain.Student{
		NISN:    req.NISN,
		Name:    req.Name,
		ClassID: req.ClassID,
	}
}

func ToUpdateStudentDomain(id int, req UpdateStudentRequest) domain.Student {
	return domain.Student{
		ID:      id,
		NISN:    req.NISN,
		Name:    req.Name,
		ClassID: req.ClassID,
	}
}
