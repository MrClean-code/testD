package model

type Deal struct {
	ID           int    `json:"id"`
	Name         string `'json:"name"`         // название услуги
	Owner        string `'json:"owner"`        // фио или компания
	Price        string `'json:"price"`        // цена
	CountReviews string `'json:"countReviews"` // количество отзывов
	Score        string `'json:"score"`        // оценка
	Link         string `'json:"link"`         // ссылка на продавца
	//TimeReceivedDeal  string `'json:"timeReceivedDeal"`  // время выполнения
	//TimeExecutionDeal string `'json:"timeExecutionDeal"` // время получения
}
