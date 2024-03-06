package videocontrol

import (
	"gin-frame/job/store"
	"log"
	"time"
)

func Count() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		<-ticker.C
		CountTriplet()
		ticker.Reset(5 * time.Second)
	}
}

func CountTriplet() {
	var video []VideoTripletData
	sql := `select gvt.video_no, 
	sum(case when gvt.is_upvote = 1 then 1 else 0 end) upvote, 
	sum(case when gvt.is_disagree = 1 then 1 else 0 end) disagree, 
	sum(case when gvt.is_coins = 1 then 1 else 0 end) coins, 
	sum(case when gvt.is_collect = 1 then 1 else 0 end) collect, 
	sum(case when gvt.is_watch = 1 then 1 else 0 end) watch 
	from gf_video_triplet gvt inner join gf_video gf on gvt.video_no = gf.video_no 
    where gf.is_del = 0 
    group by gvt.video_no`
	if err := store.GetEngine().SQL(sql).Find(&video); err != nil {
		log.Printf("find sql failed:%v", err)
	}

	for _, v := range video {
		updatesql := `update gf_video 
		set upvote = ?, disagree = ?, coins = ?, collect = ?, watch = ? 
		where video_no = ?`
		if _, err := store.GetEngine().Exec(updatesql, v.Upvote, v.Disagree, v.Coins, v.Collect, v.Watch, v.VideoNo); err != nil {
			log.Printf("update video_no[%s] triplet data failed:%v", v.VideoNo, err)
		}
	}
}
