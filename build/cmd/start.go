package cmd

import (
	"context"
	"fmt"
	"gin-frame/build/conn"
	"gin-frame/build/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
)

func start() {
	utils.FuncAction{
		{Name: "load config", Func: LoadConfig},
		{Name: "init config", Func: utils.InitConfig},
		{Name: "connection mysql", Func: conn.InitDBEngine},
		{Name: "connection redis", Func: conn.InitRedis},
		// {Name: "connection oracle", Func: conn.InitORADB},
	}.Do()
	port := viper.GetInt("http.port")
	log.Printf("web api start time:%s", time.Now().Format("2006-01-01 17:13:59"))
	r := ginInit()
	server := &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: r}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("listen and server http error: %v\n", err)
		}
	}()

	ShutDown(func() error { return server.Shutdown(context.Background()) })
}

func ShutDown(f func() error) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-quit
	// quit没有接受信号会阻塞，接受信号后会执行后续操作
	log.Printf("ShutDown Server ...\n")
	if err := f(); err != nil {
		log.Printf("ShutDown Server error:%v\n", err)
	}
	log.Printf("Server exiting\n")
}