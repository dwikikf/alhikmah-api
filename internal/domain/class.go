package domain

import "time"

type Class struct {
	ID        int
	Code      string
	Name      string
	Grade     int
	StartTime *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
