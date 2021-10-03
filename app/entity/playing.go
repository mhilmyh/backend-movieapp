package entity

import "gorm.io/gorm"

type Playing struct {
	gorm.Model
	Price	int64
	Start	string
	End	string
	Movie Movie
	Viewer Viewer
}