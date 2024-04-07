package repository

type Cache struct {
	Data map[string]interface{}
}

func (c *Cache) Add(key string, value interface{}) {
	c.Data[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
	value, exists := c.Data[key]
	return value, exists
}

func (c *Cache) Update(key string, value interface{}) {
	c.Data[key] = value
}

func (c *Cache) Delete(key string) {
	delete(c.Data, key)
}
