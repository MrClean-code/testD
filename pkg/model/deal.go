package model

type Deal struct {
	ID      int    `json:"id"`
	Name    string `'json:"name"`    // название услуги
	Owner   string `'json:"owner"`   // фио или компания
	Price   string `'json:"price"`   // цена
	Quality string `'json:"quality"` // отзывы (сред. оценка и их кол-во)
	//TimeReceivedDeal  string `'json:"timeReceivedDeal"`  // время выполнения
	//TimeExecutionDeal string `'json:"timeExecutionDeal"` // время получения
}
