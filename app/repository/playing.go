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

func (pr *playingRepository) Get() (p []entity.Playing, err error) {
	err = pr.db.Find(&p).Error
	return p, err
}

func (pr *playingRepository) Create(p *entity.Playing) (err error) {
	err = pr.db.Create(&p).Error
	return err
}

func (pr *playingRepository) Delete(id int) (err error) {
	err = pr.db.Delete(&entity.Playing{}, id).Error
	return err
}

func NewPlayingRepository(db *gorm.DB) PlayingRepository {
	return &playingRepository{db: db}
}