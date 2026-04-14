package dto

type CreateStudentRequest struct {
	NISN string `json:"nisn" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type UpdateStudentRequest struct {
	NISN string `json:"nisn" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type StudentResponse struct {
	ID   int    `json:"id"`
	NISN string `json:"nisn"`
	Name string `json:"name"`
}
