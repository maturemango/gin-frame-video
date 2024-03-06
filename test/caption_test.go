package test

import (
	"encoding/json"
	"fmt"
	"gin-frame/build/cmd"
	"gin-frame/build/conn"
	"gin-frame/build/utils"
	"gin-frame/webapi/model"
	"gin-frame/webapi/service/caption"
	"log"
	"testing"
	"time"
)

func init() {
	cmd.LoadConfig()
	utils.InitConfig()
	if err := conn.InitRedis(); err != nil {
		fmt.Println(fmt.Errorf("init redis err:%v", err))
	}
}

func TestVideoInput(t *testing.T) {
	// err := conn.Set(fmt.Sprintf(model.VideoInputKey, 1), "这是一个弹幕", 30*time.Second)
	// if err != nil {
	// 	t.Log(err)
	// }
	// val, err := conn.Get(fmt.Sprintf(model.VideoInputKey, 1));
	// if err != nil {
	// 	t.Log(err)
	// }
	// fmt.Println(val)
	err := conn.Set(fmt.Sprintf(model.VideoInputKey, "1", "1", time.Now().UnixNano()), "这是一个弹幕1", 60*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	// err = conn.Set(fmt.Sprintf(model.VideoInputKey, "1", "1", time.Now().Unix()), "新年快乐1", 60*time.Second)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	var key []string
	var cursor uint64
	for {
		key, cursor, err = conn.Scan(cursor, fmt.Sprintf(model.GetUserVideoInputKey, "1", "1")+"*", 0)
		if err != nil {
			t.Fatal(err)
		}

		for _, k := range key {
			val, err := conn.Get(k)
			if err != nil {
				t.Fatal(err)
			} else {
				fmt.Println(k, val)
			}
		}
		if cursor == 0 {
			break
		}
	}
}

// 测试隔离 不同的测试相互是隔离的
func TestGetVideoInputData(t *testing.T) {
	val, err := conn.Get(fmt.Sprintf(model.VideoInputKey, 1))
	if err != nil {
		t.Log(err)
	}
	fmt.Println(val)
}

// 视频某个时间点的所有弹幕
func TestNodeVideoInput(t *testing.T) {
	var cursor uint64
	var keyss []string
	sec := time.Now().UnixNano()
	conn.Set(fmt.Sprintf(model.VideoInputTimeKey, "1", 10, "1", sec), fmt.Sprintf(model.VideoInputKey, "1", 10, "1", sec), 60*time.Second)
	// conn.Set(fmt.Sprintf(model.VideoInputTimeKey, "1", 10, "2"), fmt.Sprintf(model.VideoInputKey, "1", 10, "2", sec), 60*time.Second)
	conn.Set(fmt.Sprintf(model.VideoInputKey, "1", 10, "1", sec), "龙年快乐", 60*time.Second)
	// conn.Set(fmt.Sprintf(model.VideoInputKey, "1", 10, "2", sec), "新年快乐", 60*time.Second)
	for {
		key, cur, err := conn.Scan(cursor, fmt.Sprintf(model.GetVideoInputTimeKey, "1", 10)+"*", 0)
		if err != nil {
			t.Fatalf("scan error:%v", err)
		}
		val, err := conn.MGet(key)
		if err != nil {
			t.Fatalf("mget error:%v", err)
		}
		for _, k := range val {
			if str, ok := k.(string); ok {
				keyss = append(keyss, str)
			}
		}

		if cur == 0 {
			break
		}
	}
	if strs, err := conn.MGet(keyss); err != nil {
		t.Fatalf("mget err:%v", err)
	} else {
		for _, s := range strs {
			fmt.Printf("caption:%v\n", s)
		}
	}
}

// 获取缓存中的结构体类型的数据
func TestUnmarshalCacheData(t *testing.T) {
	if err := cacheTestData(); err != nil {
		t.Fatal(err)
	}

	var cursor uint64
	keys, _, err := conn.Scan(cursor, fmt.Sprintf(model.GetVideoInputTimeKey, "1", 10)+"*", 0)
	if err != nil {
		t.Fatalf("sacn error:%v", err)
	}
	res, err := conn.MGet(keys)
	if err != nil {
		t.Fatalf("mget error:%v", err)
	}
	var data model.CacheVideoInput
	if len(res) <= 0 {
		t.Logf("data null")
	} else {
		for _, v := range res {
			if k, ok := v.(string); ok {
				str, err := conn.Get(k)
				if err != nil {
					t.Fatalf("get error:%v", err)
				}
				if err := json.Unmarshal([]byte(str), &data); err != nil {
					t.Fatalf("unmarshal error:%v", err)
				}

			}
		}
	}
	t.Logf("data caption:%v", data.Caption)
}

// 保存缓存弹幕
func TestSaveCacheInputData(t *testing.T) {
	if err := cacheTestData(); err != nil {
		t.Fatal(err)
	}

	var condition model.NodeInputData
	condition.Second = 10
	condition.VideoNo = "1"
	keyss, err := caption.GetCacheData(condition)
	if err != nil {
		t.Fatal(err)
	}
	n := len(keyss)
	data := make([]model.CacheVideoInput, n)
	if n <= 0 {
		t.Logf("no barrage")
	}
	for i, k := range keyss {
		str, err := conn.Get(k)
		if err != nil {
			t.Fatalf("get error:%v", err)
		}
		if err := json.Unmarshal([]byte(str), &data[i]); err != nil {
			t.Fatalf("unmarshal error:%v", err)
		}
		if _, err := conn.GetEngine().Insert(&data[i]); err != nil {
			log.Printf("insert data fail:%v", err)
			continue
		}
	}
	t.Logf("test ending")
}

func cacheTestData() error {
	sec := time.Now()
	var cache model.CacheVideoInput
	cache.VideoNo = "1"
	cache.UserId = 1
	cache.Caption = "新年快乐"
	cache.SendTime = sec
	cache.Second = 10

	b, _ := json.Marshal(cache)
	err := conn.Set(fmt.Sprintf(model.VideoInputTimeKey, "1", 10, "1", sec),
		fmt.Sprintf(model.VideoInputKey, "1", 10, "1", sec), 60*time.Second)
	if err != nil {
		return fmt.Errorf("set error:%v", err)
	}
	err = conn.Set(fmt.Sprintf(model.VideoInputKey, "1", 10, "1", sec), b, 60*time.Second)
	if err != nil {
		return fmt.Errorf("set err:%v", err)
	}
	return nil
}