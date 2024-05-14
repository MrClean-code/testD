package repository

import (
	"context"
	"fmt"
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

func (d *DealPostgres) InsertDeals(sl []model.Deal, err error) string {
	ctx := context.Background()
	tx, err := d.db.Begin(ctx)
	if err != nil {
		return ""
	}
	defer tx.Rollback(ctx)

	for _, deal := range sl {
		var ord int
		createDealQuery := fmt.Sprintf(`
		INSERT INTO %s (name, owner, price,
		                count_reviews, score, link)
		VALUES
		($1, $2, $3, $4, $5, $6)
		RETURNING id`, "deal")

		row := tx.QueryRow(ctx, createDealQuery, deal.Name, deal.Owner,
			deal.Price, deal.CountReviews, deal.Score, deal.Link)

		if err := row.Scan(&ord); err != nil {
			return ""
		}

		if err := tx.Commit(ctx); err != nil {
			return ""
		}
	}

	return ""
}
