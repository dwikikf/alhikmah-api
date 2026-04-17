package dto

import "github.com/dwikikf/alhikmah-api/internal/domain"

func ToStudentListResponse(students []domain.Student) []StudentResponse {
	var res []StudentResponse
	for _, s := range students {
		res = append(res, StudentResponse{
			ID:   s.ID,
			NISN: s.NISN,
			Name: s.Name,
		})
	}
	return res
}

func ToStudentResponse(students domain.Student) StudentResponse {
	return StudentResponse{
		ID:   students.ID,
		NISN: students.NISN,
		Name: students.Name,
	}
}

func ToCreateStudentDomain(req CreateStudentRequest) domain.Student {
	return domain.Student{
		NISN: req.NISN,
		Name: req.Name,
	}
}

func ToUpdateStudentDomain(id int, req UpdateStudentRequest) domain.Student {
	return domain.Student{
		ID:   id,
		NISN: req.NISN,
		Name: req.Name,
	}
}
