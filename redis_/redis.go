package redis_

import "C"
import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/shenqingawei/framework/nacos"

	"log"
	"time"
)

func connectionRedis(ctx context.Context, fuc func(ctx context.Context, r *redis.Client)) error {
	var err error
	client := redis.NewClient(&redis.Options{Addr: fmt.Sprintf("%v:%v", nacos.ConfigPz.Redis.Host,
		nacos.ConfigPz.Redis.Port)}) //todo:连接 redis

	fuc(ctx, client) //todo:闭包处理业务

	defer func(*redis.Client) { // todo:关闭redis
		err := client.Close()
		log.Println(err)
	}(client)

	return err
}
func IsHave(ctx context.Context, key string) (bool, error) {
	var err error
	var c int64
	err = connectionRedis(ctx, func(ctx context.Context, r *redis.Client) {
		c, err = r.Exists(ctx, key).Result()
	})
	if err != nil {
		return false, err
	}
	if c > 0 {
		return true, nil
	}
	return false, err
}
func SetMessage(ctx context.Context, key, value string, duration time.Duration) error {
	var err error
	err = connectionRedis(ctx, func(ctx context.Context, r *redis.Client) {
		err = r.Set(ctx, key, value, duration).Err()
	})
	return err
}
func GetMessage(ctx context.Context, key string) (string, error) {
	var err error
	var data string
	err = connectionRedis(ctx, func(ctx context.Context, r *redis.Client) {
		data, err = r.Get(ctx, key).Result()
	})
	return data, err
}
