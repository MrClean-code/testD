package repository

import (
	"context"
	"fmt"
	"github.com/MrClean-code/testD/pkg/model"
	"github.com/jackc/pgx/v4"
	"strconv"
)

type DealPostgres struct {
	db *pgx.Conn
}

func NewDealPostgres(db *pgx.Conn) *DealPostgres {
	return &DealPostgres{
		db: db,
	}
}

func (d *DealPostgres) GetDealsByName(name string) ([]model.Deal, error) {
	var deals []model.Deal
	query := `
		SELECT deals.name, deals.owner, deals.price,
			   deals.count_reviews, deals.score, deals.link
		FROM deals
		WHERE LENGTH(SUBSTRING(deals.name FROM $1)) >= 4
	`

	rows, err := d.db.Query(context.Background(), query, name)
	if err != nil {
		fmt.Println("Ошибка выполнения запроса:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var deal model.Deal
		err := rows.Scan(
			&deal.ID,
			&deal.Name,
			&deal.Owner,
			&deal.Price,
			&deal.CountReviews,
			&deal.Score,
			&deal.Link,
		)
		if err != nil {
			return nil, err
		}

		deals = append(deals, deal)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return deals, nil
}

func (d *DealPostgres) InsertDeals(sl []model.Deal, err error) string {
	ctx := context.Background()
	tx, err := d.db.Begin(ctx)
	if err != nil {
		return ""
	}
	defer tx.Rollback(ctx)

	fmt.Println("InsertDeals ", len(sl))
	for _, deal := range sl {
		fmt.Println(deal)
	}

	var ord int
	for _, deal := range sl {

		createDealQuery := fmt.Sprintf(`
		INSERT INTO %s (name, owner, price,
		                count_reviews, score, link)
		VALUES
		($1, $2, $3, $4, $5, $6) RETURNING id`, "deals")

		row := tx.QueryRow(ctx, createDealQuery, deal.Name, deal.Owner,
			deal.Price, deal.CountReviews, deal.Score, deal.Link)

		if err := row.Scan(&ord); err != nil {
			return ""
		}

	}

	if err := tx.Commit(ctx); err != nil {
		return ""
	}

	return "Added " + strconv.Itoa(ord)
}

func (d *DealPostgres) GetAllDeals() ([]model.Deal, error) {
	var deals []model.Deal
	query := "SELECT deals.name, deals.owner, deals.price, " +
		"deals.count_reviews, deals.score, deals.link FROM deals"

	rows, err := d.db.Query(context.Background(), query)
	if err != nil {
		fmt.Println("Ошибка выполнения запроса:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var deal model.Deal
		err := rows.Scan(
			&deal.ID,
			&deal.Name,
			&deal.Owner,
			&deal.Price,
			&deal.CountReviews,
			&deal.Score,
			&deal.Link,
		)
		if err != nil {
			return nil, err
		}

		deals = append(deals, deal)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return deals, nil
}
