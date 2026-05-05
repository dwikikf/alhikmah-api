package repository

import "errors"

var (
	ErrNotFound        = errors.New("record not found")
	ErrStudentNotFound = errors.New("student not found")
	ErrClassNotFound   = errors.New("class not found")

	ErrAlreadyCheckedIn = errors.New("student allready checked in")

	ErrDuplicate  = errors.New("data allready exists")
	ErrForeignKey = errors.New("foreign key constraint violation")
)
