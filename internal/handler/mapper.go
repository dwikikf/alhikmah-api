package handler

import (
	"github.com/dwikikf/alhikmah-api/internal/domain"
	"github.com/dwikikf/alhikmah-api/internal/dto"
)

func toStudentResponse(students []domain.Student) []dto.StudentResponse {
	var res []dto.StudentResponse
	for _, s := range students {
		res = append(res, dto.StudentResponse{
			ID:   s.ID,
			NISN: s.NISN,
			Name: s.Name,
		})
	}
	return res
}
