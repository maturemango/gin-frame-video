package controls

import (
	"fmt"
	"gin-frame/build/conn"
	"gin-frame/webapi/handlers"
	"gin-frame/webapi/model"
	_ "log"
	"sync"

	"github.com/gin-gonic/gin"
)

// 视频的播放以及观看时长的计算由前端完成
func WatchAndCountVideo(c *gin.Context) {
	var data model.UserWatchCountVideo
	if err := c.BindJSON(&data); err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	}

	var (
		wg  sync.WaitGroup
		utc model.UserTripletControls
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		if data.Duration >= 180 {
			if ok, err := conn.GetEngine().SQL(model.QueryTripletSQL, data.VideoNo, handlers.Identity()).Get(&utc); err != nil {
				// log.Printf("WatchAndCountVideo:query user[%v] triplet data failed:%v\n", handlers.Identity(), err)
				handlers.Base.Fail(c, 500, fmt.Errorf("query data failed:%v", err))
				return
			} else if ok {
				if utc.IsWatch == 1 {
					return
				}
				utc.IsWatch = 1
				if _, err := conn.GetEngine().Cols("is_watch").Where("video_no = ? and user_id = ?", data.VideoNo, handlers.Identity()).Update(&utc); err != nil {
					handlers.Base.Fail(c, 500, fmt.Errorf("update data failed:%v", err))
					return
				}
				// handlers.Base.OK(c, "")
				return
			} else if !ok {
				utc = model.UserTripletControls{}
				utc.VideoNo = data.VideoNo
				utc.UserId = handlers.Identity()
				utc.IsWatch = 1
				if _, err := conn.GetEngine().Insert(&utc); err != nil {
					handlers.Base.Fail(c, 500, fmt.Errorf("insert data failed:%v", err))
					return
				} else {
					// handlers.Base.OK(c, "")
					return
				}
			}
		}
	}()

	wg.Wait()
	return
}

// type: 0:点赞 1:投币 2:收藏 3:三连 4:点踩
func VideoTripletControl(c *gin.Context) {
	var data model.VideoTriplet
	if err := c.BindJSON(&data); err != nil {
		handlers.Base.Fail(c, 400, err)
		return
	}

	wg := new(sync.WaitGroup)

	switch data.Type {
	case 0:
		goVideo(c, wg, data, data.Type)
	case 1:
		goVideo(c, wg, data, data.Type)
	case 2:
		goVideo(c, wg, data, data.Type)
	case 3:
		goVideo(c, wg, data, data.Type)
	case 4:
		goVideo(c, wg, data, data.Type)
	default:
		handlers.Base.Fail(c, 400, fmt.Errorf("controls invalid"))
		return
	}
	wg.Wait()
}

func goVideo(c *gin.Context, wg *sync.WaitGroup, data model.VideoTriplet, typ int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := videoControlsFunc(data, typ); err != nil {
			handlers.Base.Fail(c, 500, err)
			return
		} else {
			handlers.Base.OK(c, "controls success")
			return
		}
	}()
}

func videoLikeControl(vt model.VideoTriplet) error {
	var data model.UserTripletControls
	if vt.Controller == 1 {
		data.IsUpvote = 0
		if _, err := conn.GetEngine().Cols("is_upvote").Where("video_no = ? and user_id = ?", vt.VideoNo, handlers.Identity()).Update(&data); err != nil {
			return err
		}
		return nil
	}
	ok, err := conn.GetEngine().SQL(model.QueryTripletSQL, vt.VideoNo, handlers.Identity()).Get(&data)
	if err != nil {
		return err
	} else if ok && data.IsUpvote == 1 {
		return fmt.Errorf("like exist")
	} else if ok && data.IsDisagree == 1 {
		data.IsDisagree = 0
		data.IsUpvote = 1
		if _, err := conn.GetEngine().Cols("is_upvote", "is_disagree").Where("video_no = ? and user_id = ?", vt.VideoNo, handlers.Identity()).Update(&data); err != nil {
			return fmt.Errorf("like update and disagree cancel control failed:%v", err)
		}
		return nil
	} else if ok {
		data.IsUpvote = 1
		if _, err := conn.GetEngine().Cols("is_upvote").Where("video_no = ? and user_id = ?", vt.VideoNo, handlers.Identity()).Update(&data); err != nil {
			return fmt.Errorf("like update control failed:%v", err)
		}
		return nil
	}
	data = model.UserTripletControls{}
	data.VideoNo = vt.VideoNo
	data.UserId = handlers.Identity()
	data.IsUpvote = 1
	if _, err := conn.GetEngine().Insert(&data); err != nil {
		return fmt.Errorf("like insert control failed:%v", err)
	}
	return nil
}

func videoCoinsControl(vt model.VideoTriplet) error {
	var data model.UserTripletControls
	ok, err := conn.GetEngine().SQL(model.QueryTripletSQL, vt.VideoNo, handlers.Identity()).Get(&data)
	if err != nil {
		return err
	} else if ok && data.IsCoins == 1 {
		return fmt.Errorf("coins exist")
	} else if ok {
		data.IsCoins = 1
		if _, err := conn.GetEngine().Cols("is_coins").Where("video_no = ? and user_id = ?", vt.VideoNo, handlers.Identity()).Update(&data); err != nil {
			return fmt.Errorf("coins update control failed:%v", err)
		}
		return nil
	}
	data = model.UserTripletControls{}
	data.VideoNo = vt.VideoNo
	data.UserId = handlers.Identity()
	data.IsCoins = 1
	if _, err := conn.GetEngine().Insert(&data); err != nil {
		return fmt.Errorf("coins insert control failed:%v", err)
	}
	return nil
}

func videoCollectControl(vt model.VideoTriplet) error {
	var data model.UserTripletControls
	if vt.Controller == 1 {
		data.IsCollect = 0
		if _, err := conn.GetEngine().Cols("is_collect").Where("video_no = ? and user_id = ?", vt.VideoNo, handlers.Identity()).Update(&data); err != nil {
			return err
		}
		return nil
	}
	ok, err := conn.GetEngine().SQL(model.QueryTripletSQL, vt.VideoNo, handlers.Identity()).Get(&data)
	if err != nil {
		return err
	} else if ok && data.IsCollect == 1 {
		return fmt.Errorf("collect exist")
	} else if ok {
		data.IsCollect = 1
		if _, err := conn.GetEngine().Cols("is_collect").Where("video_no = ? and user_id = ?", vt.VideoNo, handlers.Identity()).Update(&data); err != nil {
			return fmt.Errorf("collect update control failed:%v", err)
		}
		return nil
	}
	data = model.UserTripletControls{}
	data.VideoNo = vt.VideoNo
	data.UserId = handlers.Identity()
	data.IsCollect = 1
	if _, err := conn.GetEngine().Insert(&data); err != nil {
		return fmt.Errorf("collect insert control failed:%v", err)
	}
	return nil
}

func videoTripletControl(vt model.VideoTriplet) error {
	var data model.UserTripletControls
	ok, err := conn.GetEngine().SQL(model.QueryTripletSQL, vt.VideoNo, handlers.Identity()).Get(&data)
	if err != nil {
		return err
	} else if ok && data.IsUpvote == 1 && data.IsCoins == 1 && data.IsCollect == 1 {
		return fmt.Errorf("triplet exist")
	} else if ok {
		data.IsUpvote = 1
		data.IsDisagree = 0
		data.IsCoins = 1
		data.IsCollect = 1
		if _, err := conn.GetEngine().Cols("is_upvote", "is_disagree", "is_coins", "is_collect").Where("video_no = ? and user_id = ?", vt.VideoNo, handlers.Identity()).Update(&data); err != nil {
			return fmt.Errorf("triplet update control failed:%v", err)
		}
		return nil
	}
	data = model.UserTripletControls{}
	data.VideoNo = vt.VideoNo
	data.UserId = handlers.Identity()
	data.IsUpvote = 1
	data.IsCoins = 1
	data.IsCollect = 1
	if _, err := conn.GetEngine().Insert(&data); err != nil {
		return fmt.Errorf("triplet insert control failed:%v", err)
	}
	return nil
}

func videoDisagreeControl(vt model.VideoTriplet) error {
	var data model.UserTripletControls
	if vt.Controller == 1 {
		data.IsDisagree = 0
		if _, err := conn.GetEngine().Cols("is_disagree").Where("video_no = ? and user_id = ?", vt.VideoNo, handlers.Identity()).Update(&data); err != nil {
			return err
		}
		return nil
	}
	ok, err := conn.GetEngine().SQL(model.QueryTripletSQL, vt.VideoNo, handlers.Identity()).Get(&data)
	if err != nil {
		return err
	} else if ok && data.IsDisagree == 1 {
		return fmt.Errorf("disagree exist")
	} else if ok {
		data.IsUpvote = 0
		data.IsDisagree = 1
		if _, err := conn.GetEngine().Cols("is_upvote", "is_disagree").Where("video_no = ? and user_id = ?", vt.VideoNo, handlers.Identity()).Update(&data); err != nil {
			return fmt.Errorf("disagree update control failed:%v", err)
		}
		return nil
	}
	data = model.UserTripletControls{}
	data.VideoNo = vt.VideoNo
	data.UserId = handlers.Identity()
	data.IsDisagree = 1
	if _, err := conn.GetEngine().Insert(&data); err != nil {
		return fmt.Errorf("disagree insert control failed:%v", err)
	}
	return nil
}

func videoControlsFunc(data model.VideoTriplet, typ int) error {
	switch typ {
	case 0:
		return videoLikeControl(data)
	case 1:
		return videoCoinsControl(data)
	case 2:
		return videoCollectControl(data)
	case 3:
		return videoTripletControl(data)
	case 4:
		return videoDisagreeControl(data)
	default:
		return fmt.Errorf("invalid func")
	}
}
