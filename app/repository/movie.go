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
	return m, nil
}

func (mr *movieRepository) Create(m *entity.Movie) (err error) {
	return nil
}

func (mr *movieRepository) Update(m *entity.Movie) (err error) {
	return nil	
}

func (mr *movieRepository) Delete(id int) (err error) {
	return nil
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{db: db}
}
