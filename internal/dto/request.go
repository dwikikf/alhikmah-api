package dto

type CreateStudentRequest struct {
	NISN string `json:"nisn" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type UpdateStudentRequest struct {
	NISN string `json:"nisn" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type CreateClassRequest struct {
	Code  string `json:"code" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Grade string `json:"grade" binding:"required"`
}

type UpdateClassRequest struct {
	Code  string `json:"code" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Grade string `json:"grade" binding:"required"`
}
