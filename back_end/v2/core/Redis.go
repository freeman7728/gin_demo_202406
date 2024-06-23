package core

import (
	"database_lesson/global"
	"github.com/go-redis/redis"
)

func RedisInit() *redis.Client {
	config := &redis.Options{
		Addr:         global.Config.Redis.Addr,
		Password:     global.Config.Redis.Password,
		DB:           global.Config.Redis.DB, // 使用默认DB
		PoolSize:     global.Config.Redis.PoolSize,
		MinIdleConns: global.Config.Redis.MinIdleConns, //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。
		//超时
		//DialTimeout:  5 * time.Second, //连接建立超时时间，默认5秒。
		//ReadTimeout:  3 * time.Second, //读超时，默认3秒， -1表示取消读超时
		//WriteTimeout: 3 * time.Second, //写超时，默认等于读超时
		//PoolTimeout:  4 * time.Second, //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。
	}
	return redis.NewClient(config)
}
