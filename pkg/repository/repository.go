package repository

import (
	"github.com/MrClean-code/testD/pkg/model"
	"github.com/jackc/pgx/v4"
)

type DealList interface {
	GetDealsByName(string) ([]model.Deal, error)
	GetAllDeals() ([]model.Deal, error)
	InsertDeals([]model.Deal, error) string
}

type Repository struct {
	DealList
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		DealList: NewDealPostgres(db),
	}
}
