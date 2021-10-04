package server

import (
	"app/db"
	"app/entity"

	"gorm.io/gorm"
)

func NewDBConnection() *gorm.DB {
	d := db.NewPostgreDatabase()

	d.AutoMigrate(&entity.Movie{})
	d.AutoMigrate(&entity.Playing{})
	d.AutoMigrate(&entity.Viewer{})

	return d
}