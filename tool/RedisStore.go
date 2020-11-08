package tool

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/mojocn/base64Captcha"
	"log"
	"time"
)

type RedisStore struct {
	client *redis.Client
	ctx    context.Context
}

var RediStore RedisStore

func InitRedisStore() *RedisStore {
	config := GetConfig().RedisConfig
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr + ":" + config.Port,
		Password: config.Password,
		DB:       config.Db,
	})

	ctx := context.Background()

	RediStore = RedisStore{
		client: client,
		ctx:    ctx,
	}

	//
	base64Captcha.SetCustomStore(&RediStore)

	return &RediStore
}

func (rs *RedisStore) Set(id string, value string) {
	err := rs.client.Set(rs.ctx, id, value, time.Minute*10).Err()
	if err != nil {
		log.Println(err)
	}
}

func (rs *RedisStore) Get(id string, clear bool) string {
	val, err := rs.client.Get(rs.ctx, id).Result()
	if err != nil {
		log.Println(err)
		return ""
	}
	if clear {
		err := rs.client.Del(rs.ctx, id).Err()
		if err != nil {
			log.Println(err)
			return ""
		}
	}
	return val
}
