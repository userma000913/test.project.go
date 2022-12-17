package controller

import "hertz_demo/service"

var (
	svc *service.Service
)

func Init(s *service.Service) {
	svc = s
}
