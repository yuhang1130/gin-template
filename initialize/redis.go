package initialize

import (
	"context"
	"fmt"
	"gin-template/common/enum"
	"gin-template/global"

	"github.com/boj/redistore"
	redis "github.com/go-redis/redis/v8"
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

func initRedisStore() *redistore.RediStore {
	redisOpt := global.Config.Redis
	store, err := redistore.NewRediStoreWithDB(
		10, "tcp", fmt.Sprintf("%s:%s", redisOpt.Host, redisOpt.Port), redisOpt.Password, redisOpt.StoreDb, []byte(global.Config.Jwt.Secret),
	)
	if err != nil {
		global.Log.Errorf("redisStore Init Fail. Msg: %+v \n", err.Error())
		panic(err)
	}
	store.SetMaxAge(int(global.Config.Jwt.TTL))
	store.SetKeyPrefix(enum.SessionPrefix)

	return store
}
