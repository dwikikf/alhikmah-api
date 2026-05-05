package dto

import "github.com/dwikikf/alhikmah-api/internal/domain"

func ToClassListResponse(classes []domain.Class) []ClassResponse {
	var res []ClassResponse
	for _, c := range classes {
		res = append(res, ClassResponse{
			ID:        c.ID,
			Code:      c.Code,
			Name:      c.Name,
			Grade:     c.Grade,
			StartTime: c.StartTime,
		})
	}
	return res
}

func ToClassResponse(class domain.Class) *ClassResponse {
	return &ClassResponse{
		ID:        class.ID,
		Code:      class.Code,
		Name:      class.Name,
		Grade:     class.Grade,
		StartTime: class.StartTime,
	}
}

func ToCreateClassDomain(req CreateClassRequest) domain.Class {
	return domain.Class{
		Code:      req.Code,
		Name:      req.Name,
		Grade:     req.Grade,
		StartTime: req.StartTime,
	}
}

func ToUpdateClassDomain(id int, req UpdateClassRequest) domain.Class {
	return domain.Class{
		ID:        id,
		Code:      req.Code,
		Name:      req.Name,
		Grade:     req.Grade,
		StartTime: req.StartTime,
	}
}
