package domain

type Student struct {
	ID      int
	NISN    string
	Name    string
	ClassID int
	Class   Class
}
