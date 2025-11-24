package config

import (
	"ToDO/database"
	"ToDO/database/dbmodel"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	TaskRepository       dbmodel.TaskRepository
}

func New() (*Config, error) {
	config := Config{}

	databaseSession, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	database.Migrate(databaseSession)
	config.TaskRepository = dbmodel.NewTaskRepository(databaseSession)
	return &config, nil
}
