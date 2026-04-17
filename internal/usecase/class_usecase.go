package usecase

import (
	"context"

	"github.com/dwikikf/alhikmah-api/internal/domain"
	"github.com/dwikikf/alhikmah-api/internal/repository"
)

type ClassUseCase struct {
	ClassRepo repository.ClassRepository
	uow       repository.UnitOfWork
}

func NewClassUseCase(
	classRepo repository.ClassRepository,
	uow repository.UnitOfWork,
) *ClassUseCase {
	return &ClassUseCase{
		ClassRepo: classRepo,
		uow:       uow,
	}
}

func (u *ClassUseCase) GetAllClasses(ctx context.Context) ([]domain.Class, error) {
	return u.ClassRepo.FindAll(ctx)
}

func (u *ClassUseCase) GetClassByID(ctx context.Context, id int) (*domain.Class, error) {
	return u.ClassRepo.FindByID(ctx, id)
}

func (u *ClassUseCase) CreateClass(ctx context.Context, class domain.Class) (int, error) {
	var newID int

	err := u.uow.Do(ctx, func(repo repository.Repository) error {
		var err error
		newID, err = repo.Class().Create(ctx, class)
		return err
	})

	return newID, err
}

func (u *ClassUseCase) UpdateClass(ctx context.Context, class domain.Class) error {
	return u.uow.Do(ctx, func(repo repository.Repository) error {
		existing, err := repo.Class().FindByID(ctx, class.ID)
		if err != nil {
			return err
		}

		if existing == nil {
			return repository.ErrNotFound
		}

		return repo.Class().Update(ctx, class)
	})
}

func (u *ClassUseCase) DeleteClass(ctx context.Context, id int) error {
	return u.uow.Do(ctx, func(repo repository.Repository) error {
		return repo.Class().Delete(ctx, id)
	})
}
