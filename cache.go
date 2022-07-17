package cache

import (
	"time"
)

type Cache struct {
	db      map[string]string
	expires int64
}

func NewCache() Cache {
	return Cache{}
}

func (c *Cache) Get(key string) (string, bool) {
	if c.expires == 0 {
		return "", false
	}

	return c.db[key], true
}

func (c *Cache) Put(key, value string) {
	c.db[key] = value
}

func (c *Cache) Keys() []string {
	var keysMap []string
	if c.expires != 0 {
		for k, _ := range c.db {
			keysMap = append(keysMap, k)
		}
	}
	return keysMap
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
}
