package usecase

import (
	"github.com/dwikikf/alhikmah-api/internal/repository"
)

type ClassUsecase struct {
	ClassRepo repository.ClassRepository
	uow       repository.UnitOfWork
}

func NewClassUseCase(
	classRepo repository.ClassRepository,
	uow repository.UnitOfWork,
) *ClassUsecase {
	return &ClassUsecase{
		ClassRepo: classRepo,
		uow:       uow,
	}
}
