package database

import (
	"log"
	"ToDO/models/dbmodel"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&dbmodel.Task{},
	)
	log.Println("Database migrated succesfully")
}
