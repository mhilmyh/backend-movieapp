package repository

import (
	"movieapp/entity"

	"gorm.io/gorm"
)

type PlayingRepository interface {
	Get()	(p []entity.Playing, err error)
	Create(p *entity.Playing) (err error)
	Delete(id int) (err error)
}

type playingRepository struct {
	db *gorm.DB
}

func (mr *playingRepository) Get() (p []entity.Playing, err error) {
	return p, nil
}

func (mr *playingRepository) Create(p *entity.Playing) (err error) {
	return nil
}

func (mr *playingRepository) Delete(id int) (err error) {
	return nil
}

func NewPlayingRepository(db *gorm.DB) PlayingRepository {
	return &playingRepository{db: db}
}