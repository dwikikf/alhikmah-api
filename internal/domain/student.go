package domain

import "time"

type Student struct {
	ID        int
	NISN      string
	Name      string
	ClassID   int
	Class     Class
	CreatedAt time.Time
	UpdatedAt time.Time
}
