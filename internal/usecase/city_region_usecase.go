package usecase

import (
	"Address/internal/domain"
	"Address/internal/repo"
)

type CityRegionUseCase struct {
	repo *repo.CityRegionRepo
}

func NewCityRegionUseCase(r *repo.CityRegionRepo) *CityRegionUseCase {
	return &CityRegionUseCase{repo: r}
}

func (u *CityRegionUseCase) CreateRegion(region *domain.CityRegion) error {
	return u.repo.Create(region)
}
