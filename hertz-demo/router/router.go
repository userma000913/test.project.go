package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"hertz_demo/controller"
)

func InitRouter(h *server.Hertz) {
	v1 := h.Group("/v1")
	v1.GET("/test", controller.Test)
	v1.GET("/test/mgr", controller.TestMgr)
}
