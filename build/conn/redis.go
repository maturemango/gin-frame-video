package conn

import (
	"context"
	"gin-frame/build/utils"
	"time"

	"github.com/go-redis/redis/v8"
)

var _cache *redis.Client

func NewRedis(address, username, password string) *redis.Client {
	_cache = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		PoolSize: 300,
	})
	return _cache
}

func InitRedis() error {
	once.Do(func() {
		_cache = initRedis()
	})

	if _cache == nil {
		_cache = initRedis()
	}
	_, err := _cache.Ping(context.Background()).Result()
	return err
}

func initRedis() *redis.Client {
	return NewRedis(utils.Config.Redis.Addr,
		utils.Config.Redis.UserName,
		utils.Config.Redis.Password)
}

func Set(k string, v interface{}, exp time.Duration) error {
	_, err := _cache.Set(context.Background(), k, v, exp).Result()
	return err
}

func Get(k string) (string, error) {
	val, err := _cache.Get(context.Background(), k).Result()
	return val, err
}

func Del(k string) error {
	_, err := _cache.Del(context.Background(), k).Result()
	return err
}

func Scan(cursor uint64, match string, count int64) ([]string, uint64, error) {
	keys, cursor, err := _cache.Scan(context.Background(), cursor, match, count).Result()
	return keys, cursor, err
}

func MGet(keys []string) ([]interface{}, error) {
	val, err := _cache.MGet(context.Background(), keys...).Result()
	return val, err
}

func HGetAll(k string) (map[string]string, error) {
	result, err := _cache.HGetAll(context.Background(), k).Result()
	return result, err
}