package repository

import (
	"github.com/MrClean-code/testD/pkg/model"
	"github.com/jackc/pgx/v4"
)

type DealList interface {
	GetDealsByName() ([]model.Deal, error)
}

type Repository struct {
	DealList
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		DealList: NewDealPostgres(db),
	}
}
