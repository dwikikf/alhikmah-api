package di

import (
	"github.com/dwikikf/alhikmah-api/internal/usecase"
	"github.com/google/wire"
)

var usecaseSet = wire.NewSet(
	usecase.NewStudentUseCase,
	usecase.NewClassUseCase,
)
