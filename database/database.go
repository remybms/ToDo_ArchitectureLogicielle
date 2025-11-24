package database

import (
	"log"
	"ToDO/database/dbmodel"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&dbmodel.Task{},
	)
	log.Println("Database migrated succesfully")
}
