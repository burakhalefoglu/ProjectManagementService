package rediscachev8

import (
	"ProjectManagementService/pkg/helper"
	"context"
	"os"

	"github.com/appneuroncompany/light-logger/clogger"
	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	Client *redis.Client
}

func RedisCacheConstructor() *redisCache {
	return &redisCache{Client: getClient()}
}

func getClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     helper.ResolvePath("REDIS_HOST", "REDIS_PORT"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})
	func() {
		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			clogger.Error(&map[string]interface{}{ // use it wherever you want
				"Redis Connection Error: ": err,
			})
		}
	}()

	return client
}

func (r *redisCache) Get(key string) (map[string]string, error) {

	result := r.Client.HGetAll(context.Background(), key)
	if result.Err() != nil {
		clogger.Error(&map[string]interface{}{ // use it wherever you want
			"cache err: ": result.Err(),
		})
		return nil, result.Err()
	}
	return result.Val(), nil
}

func (r *redisCache) Set(key string, value map[string]interface{}) (success bool, err error) {
	result := r.Client.HMSet(context.Background(), key, value)
	if result.Err() != nil {
		clogger.Error(&map[string]interface{}{ // use it wherever you want
			"cache err: ": result.Err(),
		})
		return false, result.Err()
	}
	return true, nil
}

func (r *redisCache) Delete(key string, fields ...string) (success bool, err error) {
	result := r.Client.HDel(context.Background(), key, fields...)
	if result.Err() != nil {
		clogger.Error(&map[string]interface{}{ // use it wherever you want
			"cache err: ": result.Err(),
		})
		return false, result.Err()
	}
	return true, nil
}

func (r *redisCache) DeleteAll(key string) (success bool, err error) {
	result := r.Client.Del(context.Background(), key)
	if result.Err() != nil {
		clogger.Error(&map[string]interface{}{ // use it wherever you want
			"cache err: ": result.Err(),
		})
		return false, result.Err()
	}
	return true, nil
}
