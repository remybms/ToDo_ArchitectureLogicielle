package dbmodel

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	EndDate     string `json:"end-date"`
	Categorie   string `json:"categorie"`
}

type TaskRepository interface {
	Create(newTask *Task) (*Task, error)
	FindAll() ([]*Task, error)
	Delete(taskToDelete *Task) error
	Update(taskToUpdate *Task, taskId string) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *Task) (*Task, error) {
	if err := r.db.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *taskRepository) Delete(taskToDelete *Task) error {
	if err := r.db.Delete(taskToDelete).Error; err != nil {
		return err
	}
	return nil
}

func (r *taskRepository) FindAll() ([]*Task, error) {
	var tasks []*Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) Update(taskToUpdate *Task, taskId string) error {
	if err := r.db.Where("id = ?", taskId).Updates(taskToUpdate).Error; err != nil {
		return err
	}
	return nil
}
