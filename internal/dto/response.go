package dto

type Response[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    *T     `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

type StudentResponse struct {
	ID   int    `json:"id"`
	NISN string `json:"nisn"`
	Name string `json:"name"`
	// ClassID int           `json:"class_id"`
	Class ClassResponse `json:"class"`
}

type ClassResponse struct {
	ID    int    `json:"id"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Grade string `json:"grade"`
}
