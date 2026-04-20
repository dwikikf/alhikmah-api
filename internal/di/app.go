package di

import "github.com/dwikikf/alhikmah-api/internal/handler"

type App struct {
	Handlers *Handlers
}

type Handlers struct {
	Student *handler.StudentHandler
	Class   *handler.ClassHandler
}
