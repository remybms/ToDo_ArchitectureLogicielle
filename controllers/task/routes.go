package task

import (
	"ToDO/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	TaskConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Post("/", TaskConfigurator.addTaskHandler)
	router.Get("/", TaskConfigurator.tasksHandler)
	router.Put("/edit/{id}", TaskConfigurator.editTaskHandler)
	//router.Delete("/delete/{id}", TaskConfigurator.deleteTaskHandler)
	return router
}
