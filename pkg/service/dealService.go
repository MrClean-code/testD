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

func (d *DealListPostgres) GetDealsByName() ([]model.Deal, error) {
	return nil, nil
}
