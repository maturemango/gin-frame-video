package caption

import (
	"encoding/json"
	"fmt"
	"gin-frame/build/conn"
	"gin-frame/webapi/handlers"
	"gin-frame/webapi/model"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// 输入弹幕(可并发)
func InputWord(c *gin.Context) {
	var caption model.VideoInput
	if err := c.BindJSON(&caption); err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	}
	s := time.Now()
	sec := s.UnixNano()
	byt, err := cacheData(caption, handlers.Identity(), s)
	if err != nil {
		handlers.Base.Fail(c, 500, err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		err := conn.Set(fmt.Sprintf(model.VideoInputTimeKey, caption.VideoNo, caption.Sec, fmt.Sprint(handlers.Identity()), sec),
			fmt.Sprintf(model.VideoInputKey, caption.VideoNo, caption.Sec, fmt.Sprint(handlers.Identity()), sec), 60*time.Second)
		if err != nil {
			handlers.Base.Fail(c, 500, err)
			return
		}
		err = conn.Set(fmt.Sprintf(model.VideoInputKey, caption.VideoNo, caption.Sec, fmt.Sprint(handlers.Identity()), sec),
			byt, 60*time.Second)
		if err != nil {
			handlers.Base.Fail(c, 500, err)
			return
		}
	}()

	wg.Wait()
	handlers.Base.OK(c, nil)
}

func cacheData(caption model.VideoInput, id int64, sec time.Time) ([]byte, error) {
	var data model.CacheVideoInput
	data.VideoNo = caption.VideoNo
	data.UserId = id
	data.SendTime = sec
	data.Caption = caption.Caption
	data.Second = caption.Sec

	if byt, err := json.Marshal(data); err != nil {
		return nil, err
	} else {
		return byt, nil
	}
}

type captionData struct {
	Key     string `json:"key"`
	Caption string `json:"caption"`
}

// 弃用
func GetInputWord(c *gin.Context) {
	var cursor uint64
	var data []captionData
	for {
		key, cursor, err := conn.Scan(cursor, fmt.Sprintf(model.GetUserVideoInputKey, "1", fmt.Sprint(handlers.Identity()))+"*", 0)
		if err != nil {
			handlers.Base.Fail(c, 500, err)
			return
		}
		n := len(key)
		data = make([]captionData, n)
		for i, k := range key {
			data[i].Key = k
			data[i].Caption, err = conn.Get(k)
			if err != nil {
				handlers.Base.Fail(c, 500, err)
				return
			}
		}

		if cursor == 0 {
			break
		}
	}
	handlers.Base.OK(c, data)
}

type nodeData struct {
	UserId  int64  `json:"userId"`
	Caption string `json:"caption"`
}

func GetNodeInputData(c *gin.Context) {
	var condition model.NodeInputData
	var data []nodeData
	if err := c.BindJSON(&condition); err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	}

	var keyss []string
	keyss, err := GetCacheData(condition)
	if err != nil {
		handlers.Base.Fail(c, 500, err)
		return
	}
	// var err error
	n := len(keyss)
	data = make([]nodeData, n)
	if n <= 0 {
		handlers.Base.OK(c, "no barrage")
		return
	}
	// if strs, err := conn.MGet(keyss); err != nil {
	// 	handlers.Base.Fail(c, 500, fmt.Errorf("mget err:%v", err))
	// 	return
	// } else {
	// 	for i, s := range strs {
	// 		if str, ok := s.(string); ok {
	// 			data[i].Caption = str
	// 		}
	// 		// data[i].Caption = fmt.Sprint(s)
	// 	}
	// }

	for i, k := range keyss {
		str, err := conn.Get(k)
		if err != nil {
			handlers.Base.Fail(c, 500, fmt.Errorf("get error:%v", err))
			return
		}
		if err := json.Unmarshal([]byte(str), &data[i]); err != nil {
			handlers.Base.Fail(c, 500, fmt.Errorf("unmarshal error:%v", err))
			return
		}
	}

	handlers.Base.OK(c, data)
}

// 测试用接口,不做api使用
func SaveCacheInputData(c *gin.Context) {
	var condition model.NodeInputData
	if err := c.BindJSON(&condition); err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	}

	var keyss []string
	keyss, err := GetCacheData(condition)
	if err != nil {
		handlers.Base.Fail(c, 500, err)
		return
	}
	n := len(keyss)
	data := make([]model.CacheVideoInput, n)
	if n <= 0 {
		handlers.Base.OK(c, "no barrage")
		return
	}
	for i, k := range keyss {
		str, err := conn.Get(k)
		if err != nil {
			handlers.Base.Fail(c, 500, fmt.Errorf("get error:%v", err))
			return
		}
		if err := json.Unmarshal([]byte(str), &data[i]); err != nil {
			handlers.Base.Fail(c, 500, fmt.Errorf("unmarshal error:%v", err))
			return
		}
		if _, err := conn.GetEngine().Insert(&data[i]); err != nil {
			log.Printf("insert data fail:%v", err)
			continue
		}
	}

	handlers.Base.OK(c, "data save success")
}

func GetCacheData(condition model.NodeInputData) ([]string, error) {
	var keyss []string
	var cursor uint64
	for {
		key, cur, err := conn.Scan(cursor, fmt.Sprintf(model.GetVideoInputTimeKey, condition.VideoNo, condition.Second)+"*", 0)
		if err != nil {
			return keyss, fmt.Errorf("scan error:%v", err)
		}
		if len(key) <= 0 {
			break
		}
		keys, err := conn.MGet(key)
		if err != nil {
			return keyss, fmt.Errorf("mget error:%v", err)
		}
		for _, k := range keys {
			if str, ok := k.(string); ok {
				keyss = append(keyss, str)
			}
		}

		if cur == 0 {
			break
		}
	}
	return keyss, nil
}
