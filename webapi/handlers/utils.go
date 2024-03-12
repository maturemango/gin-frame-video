package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
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

func EncodeCrypto(code string) (string, error) {
	p := []byte(code)
	h := sha256.New()
	if _, err := h.Write(p); err != nil {
		return "", errors.New(err.Error())
	}
	hashed := hex.EncodeToString(h.Sum(nil))
	return hashed, nil
}
