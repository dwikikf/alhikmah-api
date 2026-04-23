package dto

type CreateStudentRequest struct {
	NISN    string `json:"nisn" binding:"required" validate:"required,numeric,len=10,min=10,max=10"`
	Name    string `json:"name" binding:"required" validate:"required,min=3,max=100"`
	ClassID int    `json:"class_id" binding:"required" validate:"required,numeric,min=1"`
}

type UpdateStudentRequest struct {
	NISN    string `json:"nisn" binding:"required" validate:"required,numeric,len=10,min=10,max=10"`
	Name    string `json:"name" binding:"required" validate:"required,min=3,max=100"`
	ClassID int    `json:"class_id" binding:"required" validate:"required,numeric,min=1"`
}

type CreateClassRequest struct {
	Code  string `json:"code" binding:"required" validate:"required,min=5,max=20"`
	Name  string `json:"name" binding:"required" validate:"required,min=3,max=100"`
	Grade int    `json:"grade" binding:"required" validate:"required,oneof=1 2 3 4 5 6,numeric"`
}

type UpdateClassRequest struct {
	Code  string `json:"code" binding:"required" validate:"required,min=5,max=20"`
	Name  string `json:"name" binding:"required" validate:"required,min=3,max=100"`
	Grade int    `json:"grade" binding:"required" validate:"required,oneof=1 2 3 4 5 6,numeric"`
}
