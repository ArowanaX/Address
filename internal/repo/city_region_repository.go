package repo

import (
	"Address/internal/domain"
	"gorm.io/gorm"
)

type CityRegionRepo struct {
	DB *gorm.DB
}

func NewCityRegionRepo(db *gorm.DB) *CityRegionRepo {
	return &CityRegionRepo{DB: db}
}
func (r *CityRegionRepo) Create(region *domain.CityRegion) error {
	return r.DB.Create(region).Error
}
