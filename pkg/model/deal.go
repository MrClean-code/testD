package model

type DealDto struct {
	ID           int     `json:"id" db:"id"`
	Name         string  `'json:"name" db:"name"`                  // название услуги
	Owner        string  `'json:"owner" db:"owner"`                // фио или компания
	Price        string  `'json:"price" db:"price"`                // цена
	CountReviews int     `'json:"countReviews" db:"count_reviews"` // количество отзывов
	Score        float64 `'json:"score" db:"score"`                // оценка
	Link         string  `'json:"link" db:"link"`                  // ссылка на специалиста
}

type Deal struct {
	ID           int    `json:"id" db:"id"`
	Name         string `'json:"name" db:"name"`                  // название услуги
	Owner        string `'json:"owner" db:"owner"`                // фио или компания
	Price        string `'json:"price" db:"price"`                // цена
	CountReviews string `'json:"countReviews" db:"count_reviews"` // количество отзывов
	Score        string `'json:"score" db:"score"`                // оценка
	Link         string `'json:"link" db:"link"`                  // ссылка на специалиста
}
