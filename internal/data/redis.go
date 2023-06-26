package data

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type Cache struct {
	db *redis.Client
}

func NewCache(db *redis.Client) *Cache {
	return &Cache{db: db}
}

func (c *Cache) RSet(name string, val []byte) (ret bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Println(ctx, name, val)
	err := c.db.Set(ctx, name, val, time.Hour).Err()
	if err != nil {
		fmt.Println("redis 缓存同步完成", err)
		panic(err)
	}
	return true
}
func (c *Cache) RGet(ctx context.Context, key string) (string, error) {
	return c.db.Get(ctx, key).Result()
}
