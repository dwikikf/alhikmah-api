package di

import (
	infraRepo "github.com/dwikikf/alhikmah-api/internal/infrastucture/repository"
	"github.com/dwikikf/alhikmah-api/internal/repository"
	"github.com/google/wire"
)

var repositorySet = wire.NewSet(
	//constructor
	infraRepo.NewStudentRepository,
	infraRepo.NewClassRepository,
	infraRepo.NewUnitOfWork,

	// binding interface to implementation
	wire.Bind(new(repository.StudentRepository), new(*infraRepo.StudentRepositoryImpl)),
	wire.Bind(new(repository.ClassRepository), new(*infraRepo.ClassRepositoryImpl)),
	wire.Bind(new(repository.UnitOfWork), new(*infraRepo.UnitOfWorkImpl)),
)
