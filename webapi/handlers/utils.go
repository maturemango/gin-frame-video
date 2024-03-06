package handlers

import (
	"gin-frame/build/conn"
	"gin-frame/webapi/model"
	"time"
)

func AddSystemLog(id int64, addr string, level model.LogLevel, desc model.UserServiceDesc, detail model.OperateDetail) error {
	var log model.OperateLog
	log.UserId = id
	log.Addr = addr
	log.LogLevel = string(level)
	log.OperateTime = time.Now()
	log.OperateDesc = string(desc)
	log.Detail = string(detail)
	if _, err := conn.GetEngine().InsertOne(&log); err != nil {
		return err
	}
	return nil
}