package service

import (
	"github.com/MrClean-code/testD/pkg/model"
	"github.com/MrClean-code/testD/pkg/repository"
)

type DealListPostgres struct {
	repos repository.DealList
	//cache *repository.Cache
}

func NewDealListPostgres(repos repository.DealList) *DealListPostgres {
	return &DealListPostgres{
		repos: repos,
	}
}

func (d *DealListPostgres) GetDealsByName(name string) ([]model.Deal, error) {
	return d.repos.GetDealsByName(name)
}

func (d *DealListPostgres) InsertDeals(sl []model.Deal, err error) string {
	return d.repos.InsertDeals(sl, err)
}

func (d *DealListPostgres) GetAllDeals() ([]model.Deal, error) {
	return d.repos.GetAllDeals()
}
