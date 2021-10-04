package repository

import (
	"movieapp/entity"

	"gorm.io/gorm"
)

type MovieRepository interface {
	Get()	(m []entity.Movie, err error)
	Create(m *entity.Movie) (err error)
	Update(m *entity.Movie) (err error)
	Delete(id int) (err error)
}

type movieRepository struct {
	db *gorm.DB
}

func (mr *movieRepository) Get() (m []entity.Movie, err error) {
	err = mr.db.Find(&m).Error
	return m, err
}

func (mr *movieRepository) Create(m *entity.Movie) (err error) {
	err = mr.db.Create(&m).Error
	return err
}

func (mr *movieRepository) Update(m *entity.Movie) (err error) {
	err = mr.db.Model(&m).Updates(m).Error
	return err
}

func (mr *movieRepository) Delete(id int) (err error) {
	err = mr.db.Delete(&entity.Movie{}, id).Error
	return err
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{db: db}
}
