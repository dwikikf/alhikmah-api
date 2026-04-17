package dto

import "github.com/dwikikf/alhikmah-api/internal/domain"

func ToClassListResponse(classes []domain.Class) []ClassResponse {
	var res []ClassResponse
	for _, c := range classes {
		res = append(res, ClassResponse{
			ID:    c.ID,
			Code:  c.Code,
			Name:  c.Name,
			Grade: c.Grade,
		})
	}
	return res
}

func ToClassResponse(class domain.Class) ClassResponse {
	return ClassResponse{
		ID:    class.ID,
		Code:  class.Code,
		Name:  class.Name,
		Grade: class.Grade,
	}
}
