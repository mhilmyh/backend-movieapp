package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgreDatabase() *gorm.DB  {
	dsn := "host=db user=movieuser password=moviepass dbname=moviedb port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})	
	if err != nil {
		panic(err)
	}

	return db
}