package videocontrol

import (
	"gin-frame/job/store"
	"log"
)

func init() {
	if err := store.InitDbEngine(); err != nil {
		log.Fatalf("init db failed:%v", err)
	}
	if err := store.InitRedisClient(); err != nil {
		log.Fatalf("init redis failed:%v", err)
	}

	go Count()
}