package repository

import "errors"

var (
	ErrNotFound = errors.New("record not found")

	ErrDuplicate  = errors.New("data allready exists")
	ErrForeignKey = errors.New("foreign key constraint violation")
)
