package usecase

import (
	"context"

	"github.com/dwikikf/alhikmah-api/internal/dto"
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

func (uc *ClassUsecase) GetAllClasses(ctx context.Context) ([]dto.ClassResponse, error) {
	classes, err := uc.ClassRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return dto.ToClassListResponse(classes), nil
}

func (uc *ClassUsecase) GetClassByID(ctx context.Context, id int) (*dto.ClassResponse, error) {
	class, err := uc.ClassRepo.FindByID(ctx, id)
	if err != nil {
		if class == nil {
			return nil, repository.ErrClassNotFound
		}
		return nil, err
	}

	return dto.ToClassResponse(*class), nil
}

func (uc *ClassUsecase) CreateClass(ctx context.Context, class dto.CreateClassRequest) (int, error) {
	var id int
	classDomain := dto.ToCreateClassDomain(class)

	err := uc.uow.Do(ctx, func(repo repository.Repository) error {
		var err error
		id, err = repo.Class().Create(ctx, classDomain)
		return err
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc *ClassUsecase) UpdateClass(ctx context.Context, id int, class dto.UpdateClassRequest) error {
	classDomain := dto.ToUpdateClassDomain(id, class)

	return uc.uow.Do(ctx, func(repo repository.Repository) error {

		existing, err := repo.Class().FindByID(ctx, id)
		if err != nil {
			return err
		}

		if existing == nil {
			return repository.ErrClassNotFound
		}

		return repo.Class().Update(ctx, classDomain)
	})
}

func (uc *ClassUsecase) DeleteClass(ctx context.Context, id int) error {
	return uc.uow.Do(ctx, func(repo repository.Repository) error {
		return repo.Class().Delete(ctx, id)
	})
}
