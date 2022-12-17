package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/registry/nacos"
	"log"
)

// hertz框架整合nacos做服务发现/注册
func main() {
	addr := "127.0.0.1:8888"
	r, err := nacos.NewDefaultNacosRegistry()
	if err != nil {
		log.Fatal(err)
		return
	}
	h := server.Default(
		server.WithHostPorts(addr),
		server.WithRegistry(r, &registry.Info{
			ServiceName: "hertz.test.demo",
			Addr:        utils.NewNetAddr("tcp", addr),
			Weight:      10,
			Tags:        nil,
		}),
	)
	h.POST("/ping", func(c context.Context, ctx *app.RequestContext) {
		req := struct {
			Name string `json:"name" form:"name"`
			Age  int    `json:"age" form:"age"`
		}{}
		ctx.Bind(&req)
		ctx.JSON(consts.StatusOK, utils.H{"name": req.Name, "age": req.Age})
	})
	h.GET("/t", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "ping"})
	})
	h.Spin()
}
