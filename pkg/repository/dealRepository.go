package repository

import (
	"github.com/MrClean-code/testD/pkg/model"
	"github.com/jackc/pgx/v4"
)

type DealPostgres struct {
	db *pgx.Conn
}

func NewDealPostgres(db *pgx.Conn) *DealPostgres {
	return &DealPostgres{
		db: db,
	}
}

func (d *DealPostgres) GetDealsByName() ([]model.Deal, error) {
	//sql
	return nil, nil
}
