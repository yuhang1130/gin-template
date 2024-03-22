package initialize

import (
	"context"
	"fmt"
	"gin-template/global"

	"github.com/go-redis/redis/v8"
)

func initRedis() *redis.Client {
	redisOpt := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisOpt.Host, redisOpt.Port),
		Password: redisOpt.Password, // no password set
		DB:       redisOpt.DataBase, // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Log.Errorf("redis connect ping failed, err: %s\n", err)
		panic(err)
	}
	global.Log.Infof("redis connect ping successes. pong: %s\n", pong)
	return client
}
