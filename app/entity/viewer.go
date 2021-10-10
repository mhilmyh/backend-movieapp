package entity

import "gorm.io/gorm"

type Viewer struct {
	gorm.Model
	Name	string
	PlayingID int
	Playing
}