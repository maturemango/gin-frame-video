package captionstore

type CaptionEvent struct{}

func (ce *CaptionEvent) Run() {
	SaveCacheInputData()
}

func (ce *CaptionEvent) Spec() string {
	return "0 0 8 * * ?"  // 每天早上八点执行一次定时任务
}