package task

import (
	"net/http"
	"ToDO/config"
	"ToDO/models/dbmodel"
	"ToDO/controllers/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type TaskConfigurator struct {
	*config.Config
}

func New(configuration *config.Config) *TaskConfigurator {
	return &TaskConfigurator{configuration}
}

func tasksToModel(tasks []*dbmodel.Task) []models.Task {
	tasksToModel := &models.Task{}
	tasksEdited := []models.Task{}
	for _, Task := range tasks {
		tasksToModel.Title = Task.Title
		tasksToModel.Description = Task.Description
		tasksToModel.EndDate = Task.EndDate
		tasksToModel.Categorie = Task.Categorie
		tasksEdited = append(tasksEdited, *tasksToModel)
	}
	return tasksEdited
}

func (config *TaskConfigurator) tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := config.TaskRepository.FindAll()
	tasksEdited := tasksToModel(tasks)
	if err != nil {
		render.JSON(w, r, map[string]string{"Error": "Failed to load all the tasks"})
		return
	}
	render.JSON(w, r, tasksEdited)
}

func (config *TaskConfigurator) addTaskHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.Task{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	addTask := &dbmodel.Task{Title: req.Title, Description: req.Description, EndDate: req.EndDate, Categorie: req.Categorie}
	config.TaskRepository.Create(addTask)
	render.JSON(w, r, map[string]string{"success": "New Task successfully added"})
}

func (config *TaskConfigurator) editTaskHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.Task{}
	TaskId := chi.URLParam(r, "id")
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	updatedTask := &dbmodel.Task{Title: req.Title, Description: req.Description, EndDate: req.EndDate, Categorie: req.Categorie}
	config.TaskRepository.Update(updatedTask, TaskId)
	render.JSON(w, r, map[string]string{"success": "Task successfully updated"})
}

/*func (config *TaskConfigurator) deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	TaskId := chi.URLParam(r, "id")
	if err != nil {
		render.JSON(w, r, map[string]string{"Error": "Failed to find the wanted Task"})
		return
	}
	config.TaskRepository.Delete(Task[0])
	render.JSON(w, r, map[string]string{"success": "Task successfully deleted"})
}*/
