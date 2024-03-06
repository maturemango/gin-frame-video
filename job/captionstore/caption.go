package captionstore

import (
	"encoding/json"
	"fmt"
	"gin-frame/job/store"
	"log"
)

// 定时保存视频弹幕，当前并没有视频相关的表
func SaveCacheInputData() {
	var cursor uint64
	var vals []string
	for {
		keys, cur, err := store.Scan(cursor, fmt.Sprintf(VideoInputKey, "1")+"*", 0)
		if err != nil {
			log.Fatalf("cache scan failed:%v", err)
		}
		if len(keys) <= 0 {
			break
		}

		keyss, err := store.MGet(keys)
		if err != nil {
			log.Fatalf("cache mget failed:%v", err)
		}
		for _, v := range keyss {
			if str, ok := v.(string); ok {
				vals = append(vals, str)
			}
		}

		if cur == 0 {
			break
		}
	}

	n := len(vals)
	data := make([]CacheInputData, n)
	if n <= 0 {
		log.Println("cache data nil")
	}
	for i, k := range vals {
		str, err := store.Get(k)
		if err != nil {
			log.Fatalf("cache get failed:%v", err)
			continue
		}
		if err := json.Unmarshal([]byte(str), &data[i]); err != nil {
			log.Fatalf("unmarshal failed:%v", err)
			continue
		}
		if _, err := store.GetEngine().Insert(&data[i]); err != nil {
			log.Printf("insert data failed:%v", err)
			continue
		}
	}
}