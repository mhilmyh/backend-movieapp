package repository

import (
	"movieapp/entity"

	"gorm.io/gorm"
)

type ViewerRepository interface {
	Get()	(v []entity.Viewer, err error)
	Create(v *entity.Viewer) (err error)
	Delete(id int) (err error)
}

type viewerRepository struct {
	db *gorm.DB
}

func (mr *viewerRepository) Get() (v []entity.Viewer, err error) {
	return v, nil
}

func (mr *viewerRepository) Create(v *entity.Viewer) (err error) {
	return nil
}

func (mr *viewerRepository) Delete(id int) (err error){
	return nil	
}

func NewViewerRepository(db *gorm.DB) ViewerRepository {
	return &viewerRepository{db: db}
}