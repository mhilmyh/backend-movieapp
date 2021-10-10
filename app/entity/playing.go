package entity

import (
	"time"

	"gorm.io/gorm"
)

type Playing struct {
	gorm.Model
	Price	int64
	Start	time.Time
	End	time.Time
	MovieID int
	Movie Movie
	ViewerID int
	Viewer Viewer
}