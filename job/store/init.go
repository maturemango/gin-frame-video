package store

import (
	"context"
	"gin-frame/build/conn"
	"log"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"xorm.io/xorm"
)

var (
	Iviper  *viper.Viper
	once    sync.Once
	engine  *xorm.Engine
	_cache  *redis.Client
)

func init() {
	Iviper = viper.New()
	pwd, _ := os.Executable()
	Iviper.AddConfigPath(path.Join(filepath.Dir(pwd), "conf"))
	Iviper.AddConfigPath("D:\\gopath\\wtchain.com\\gin-frame\\job\\store\\conf")
	Iviper.AddConfigPath("./conf")
	Iviper.SetConfigName("job")
	Iviper.OnConfigChange(func(in fsnotify.Event) {})
	if err := Iviper.ReadInConfig(); err != nil {
		panic(err)
	}

	Iviper.WatchConfig()

	if err := InitDbEngine(); err != nil {
		log.Fatalf("init db failed:%v", err)
	}
	if err := InitRedisClient(); err != nil {
		log.Fatalf("init redis failed:%v", err)
	}

}

func InitDbEngine() error {
	once.Do(func() {
		engine = initEngine()
	})

	engine.ShowSQL(Iviper.GetBool("mysql.showsql"))
	return engine.Ping()
}

func initEngine() *xorm.Engine {
	return conn.NewXormDB(Iviper.GetString("mysql.host"),
		Iviper.GetString("mysql.database"),
		Iviper.GetString("mysql.username"),
		Iviper.GetString("mysql.password"))
}

func GetEngine() *xorm.Engine {
	if engine == nil {
		engine = initEngine()
	}

	return engine
}

func InitRedisClient() error {
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
	return conn.NewRedis(Iviper.GetString("redis.addr"),
		Iviper.GetString("redis.username"),
		Iviper.GetString("redis.password"))
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