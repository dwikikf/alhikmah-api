package di

import (
	"github.com/dwikikf/alhikmah-api/internal/handler"
	"github.com/google/wire"
)

var handlerSet = wire.NewSet(
	handler.NewStudentHandler,
	handler.NewClassHandler,
)
