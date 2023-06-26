package account

import (
	"admin_api/internal/model"
	"admin_api/internal/response"
	"context"
	"fmt"
	"go.uber.org/zap"
)

type CacheRepo interface {
	Create(name, val string) (ret bool)
	Find(name string) (*model.Cache, error)
}

type Cache struct {
	cacheRepo CacheRepo
}

func NewCacheService(CacheRepo CacheRepo) *Cache {
	return &Cache{cacheRepo: CacheRepo}
}

func (c *Cache) Create(name, val string) bool {
	return c.cacheRepo.Create(name, val)
}
func (c *Cache) Check(name string) (*response.Cache, error) {
	cc, err := c.cacheRepo.Find(name)
	if err != nil {
		zap.S().Errorf("%+#v\n", err)
		return nil, err
	}
	return &response.Cache{Key: cc.Name, Value: cc.Value}, nil
}

type RedisRepo interface {
	RSet(name string, val []byte) (ret bool)
	RGet(ctx context.Context, key string) (string, error)
}

type Redis struct {
	redisRepo RedisRepo
}

func NewRedisService(redisRepo RedisRepo) *Redis {
	return &Redis{redisRepo: redisRepo}
}

func (r *Redis) Create(name string, val []byte) bool {
	fmt.Println("9999999999999")
	return r.redisRepo.RSet(name, val)
}
func (r *Redis) Find(ctx context.Context, key string) (string, error) {
	return r.redisRepo.RGet(ctx, key)
}
