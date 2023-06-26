package MyCache

import (
	"admin_api/global"
	"admin_api/internal/data"
	"admin_api/internal/response"
	"admin_api/internal/service/store/mysql/account"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Cache struct{}

type Req struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

func (c *Cache) Crete(ctx *gin.Context) {
	db := account.NewCacheService(data.NewCacheRepo(global.DB))
	//db.Create("key1", "mysql数据")
	rdb := account.NewRedisService(data.NewCache(global.RDB))

	var key *Req
	if err := ctx.ShouldBind(&key); err != nil {
		ctx.JSON(http.StatusOK, response.Error(1, err.Error()))
		fmt.Println("11111", key.Key, err.Error())
		return
	}
	fmt.Println("11111", key)
	ret := db.Create(key.Key, key.Value)
	if ret == true {
		ctx.JSON(http.StatusOK, gin.H{
			"database": "mysql",
			"status":   "insert success",
		})
	}

	rc, err := db.Check(key.Key)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Error(1, err.Error()))
		fmt.Println("11111", key.Key, err.Error())
		return
	}

	// 将数据缓存到Redis中
	k := "user:" + rc.Key
	val, err := json.Marshal(rc)
	if err != nil {
		panic(err)
	}
	fmt.Println("Marshal success", k, val)
	re := rdb.Create(k, val)
	if re == true {
		ctx.JSON(http.StatusOK, gin.H{
			"database": "redis",
			"status":   "insert redis success",
		})
	}
}

func (c *Cache) Find(ctx *gin.Context) {
	// 从请求中获取用户名
	name := ctx.Query("name")
	key := "user:" + name
	fmt.Println("11111111111111111111111111", name)
	// 优先从缓存中获取用户信息
	rdb := account.NewRedisService(data.NewCache(global.RDB))
	val, err := rdb.Find(ctx, key)
	if err == nil {
		// 如果缓存中存在用户信息，直接返回
		var user *response.Cache
		err := json.Unmarshal([]byte(val), &user)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"database": "redis",
			"status":   "success",
			"data":     user,
		})
		return
	}

	// 如果缓存中不存在用户信息，从数据库中获取
	var user *response.Cache
	db := account.NewCacheService(data.NewCacheRepo(global.DB))
	user, err = db.Check(name)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"database": "mysql",
			"status":   "failed",
			"reason":   err.Error(),
		})
		return
	}

	// 将用户信息存储到 Redis 缓存中
	v, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	ret := rdb.Create(key, v)
	if ret != true {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}
