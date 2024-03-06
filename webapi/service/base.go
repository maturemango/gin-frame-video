package service

import "gin-frame/webapi/handlers"

var Svc *Service

func init() {
	Svc = NewService()
}

func NewService() *Service {
	return &Service{}
}

type Service struct {
	handlers.BaseHandler
}