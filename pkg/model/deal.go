package model

type Deal struct {
	ID           int    `json:"id" db:"id"`
	Name         string `'json:"name" db:"name"`                  // название услуги
	Owner        string `'json:"owner" db:"owner"`                // фио или компания
	Price        string `'json:"price" db:"price"`                // цена
	CountReviews string `'json:"countReviews" db:"count_reviews"` // количество отзывов
	Score        string `'json:"score" db:"score"`                // оценка
	Link         string `'json:"link" db:"link"`                  // ссылка на продавца
}
