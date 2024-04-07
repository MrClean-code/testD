package service

import (
	"github.com/MrClean-code/testD/pkg/model"
	"github.com/MrClean-code/testD/pkg/repository"
)

type DealList interface {
	GetDealsByName() ([]model.Deal, error)
}

type Service struct {
	DealList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		DealList: NewDealListPostgres(repos.DealList),
	}
}
