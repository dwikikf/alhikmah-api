package dto

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type StudentResponse struct {
	ID   int    `json:"id"`
	NISN string `json:"nisn"`
	Name string `json:"name"`
}

type ClassResponse struct {
	ID    int    `json:"id"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Grade string `json:"grade"`
}
