package repository

import (
	"app/cache"
	"app/entity"

	"gorm.io/gorm"
)

type ViewerRepository interface {
	Get()	(v []entity.Viewer, err error)
	Create(v *entity.Viewer) (err error)
	Delete(id int) (err error)
}

type viewerRepository struct {
	db *gorm.DB
	cache *cache.Redis
}

func (vr *viewerRepository) Get() (v []entity.Viewer, err error) {
	err = vr.db.Find(&v).Error
	return v, err
}

func (vr *viewerRepository) Create(v *entity.Viewer) (err error) {
	err = vr.db.Create(&v).Error
	return err
}

func (vr *viewerRepository) Delete(id int) (err error){
	err = vr.db.Delete(&entity.Viewer{}, id).Error
	return err	
}

func NewViewerRepository(db *gorm.DB, c *cache.Redis) ViewerRepository {
	return &viewerRepository{db: db, cache: c}
}