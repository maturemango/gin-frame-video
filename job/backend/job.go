package backend

import (
	"gin-frame/job/captionstore"
	"gin-frame/job/store"
	"log"

	"github.com/robfig/cron/v3"
)

var TaskMap = make(map[string]Task)

type Task interface {
	cron.Job
	Spec() string
}

type BackgroundJob struct {
	c     *cron.Cron
	done chan struct{}
}

func (bj *BackgroundJob) Run() {
	bj.c = cron.New(
		cron.WithParser(cron.NewParser(
			cron.SecondOptional|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow|cron.Descriptor)), 
			// 解析器支持秒（可选的）、分钟、小时、日期、月份、星期和描述符
		cron.WithLogger(cron.VerbosePrintfLogger(log.Default())),
		cron.WithChain(cron.Recover(cron.DefaultLogger)),
	)

	bj.addJobs()
	bj.c.Start()
	bj.done = make(chan struct{})
	<-bj.done
}

func (bj *BackgroundJob) addJobs() {
	// 根据配置文件选择需要执行的job
	var tasks []Task
	if store.Iviper.GetBool("job.captionstore") {
		tasks = append(tasks, &captionstore.CaptionEvent{})
	}
	if err := bj.AddTask(tasks...); err != nil {
		log.Fatalf("add job error:%v", err)
	}
	
	// 执行添加的所有的job
	// if err := bj.AddTask(
	// 	&captionstore.CaptionEvent{},
	// ); err != nil {
	// 	log.Fatalf("add job error:%v", err)
	// }
}

func (bj *BackgroundJob) AddTask(t ...Task) error {
	for _, v := range t {
		_, err := bj.c.AddJob(v.Spec(), v)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func (bj *BackgroundJob) Stop() {
	if bj.c != nil {
		bj.c.Stop()
	}
	bj.done <- struct{}{}
}