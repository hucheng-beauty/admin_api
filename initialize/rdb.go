package initialize

import (
	"admin_api/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func RedisDB() {
	global.RDB = RedisClient()
}

func RedisClient() *redis.Client {
	// 从config文件中读取Redis连接信息
	redisInfo := global.ServerConfig.RedisInfo

	// 创建Redis客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisInfo.Endpoint, // Redis服务器地址和端口号
		Password: redisInfo.Password, // Redis服务器密码
		DB:       redisInfo.Database, // Redis数据库编号
	})

	// 测试连接
	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis连接测试结果:", pong)

	return rdb
}
