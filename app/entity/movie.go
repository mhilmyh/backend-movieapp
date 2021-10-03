package entity

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title				string
	Desc				string
	Rating			int8
	Author			string
}