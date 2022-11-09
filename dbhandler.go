package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	ctx             = context.Background()
	linkRedisMethod sync.Once
	DbRedis         *redis.Client
)

// 连接redis数据库
// 且只会执行一次
func initRedis() {
	linkRedisMethod.Do(func() {
		//连接数据库
		DbRedis = redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379", // 对应的ip以及端口号
			Password: "",               // 数据库的密码
			DB:       0,                // 数据库的编号，默认的话是0
		})
		// 连接测活
		_, err := DbRedis.Ping(ctx).Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("连接Redis成功")
	})
}

// 登录成功时向数据库中插入相关信息
func InsertMapData(service string, uid string, username string) (err error) {
	loginInfo := make(map[string]string)
	loginInfo["service"] = service
	loginInfo["username"] = username
	loginInfo["uid"] = uid
	err = DbRedis.HSet(ctx, uid, loginInfo).Err()
	return
}

// 删除登录状态
func RemoveLoginHistory(uid string) (err error) {
	err = DbRedis.Del(ctx, uid).Err()
	return
}

// 检测是否已经登录了
func CheckLoadingState(uid string) (err error) {
	_, err = DbRedis.HGetAll(ctx, uid).Result()
	return
}
