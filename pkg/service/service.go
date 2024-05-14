package service

import (
	"github.com/MrClean-code/testD/pkg/model"
	"github.com/MrClean-code/testD/pkg/repository"
)

type DealList interface {
	GetDealsByName() ([]model.Deal, error)
	InsertDeals([]model.Deal, error) string
}

type Service struct {
	DealList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		DealList: NewDealListPostgres(repos.DealList),
	}
}

//type DealListPostgres struct {
//	repos repository.DealList
//	//cache *repository.Cache
//}
//
//func NewDealListPostgres(repos repository.DealList) *DealListPostgres {
//	return &DealListPostgres{
//		repos: repos,
//	}
//}
//
//func (d *DealListPostgres) GetDealsByName() ([]model.Deal, error) {
//	return nil, nil
//}
