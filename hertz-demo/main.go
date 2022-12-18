package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	hertzSentinel "github.com/hertz-contrib/opensergo/sentinel/adapter"
	"github.com/hertz-contrib/registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/utils"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"hertz_demo/conf"
	"hertz_demo/controller"
	"hertz_demo/router"
	"hertz_demo/service"
	"log"
)

var wg sync.WaitGroup

type Test struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {

	c := conf.InitConfig()
	s := service.New(c)
	controller.Init(s)

	// 服务注册
	//r, err := nacos.NewDefaultNacosRegistry()
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848),
	}

	cc := constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "info",
	}

	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	wg.Add(2)
	go func() {
		defer wg.Done()
		addr := "127.0.0.1:8888"
		r := nacos.NewNacosRegistry(cli)
		h := server.Default(
			server.WithHostPorts(addr),
			server.WithRegistry(r, &registry.Info{
				ServiceName: "demo1",
				Addr:        utils.NewNetAddr("tcp", addr),
				Weight:      10,
				Tags:        nil,
			}),
		)
		h.POST("/ping", func(c context.Context, ctx *app.RequestContext) {
			t := Test{}
			if err := ctx.Bind(&t); err != nil {
				ctx.String(consts.StatusOK, err.Error())
				return
			}
			ctx.JSON(consts.StatusOK, t)
		})

		h.Spin()
	}()

	go func() {
		defer wg.Done()
		addr := fmt.Sprintf("127.0.0.1:%d", c.Port)
		r := nacos.NewNacosRegistry(cli)
		tracer, cfg := hertztracing.NewServerTracer()
		h := server.Default(
			server.WithHostPorts(addr),
			server.WithRegistry(r, &registry.Info{
				ServiceName: c.Name,
				Addr:        utils.NewNetAddr("tcp", addr),
				Weight:      10,
				Tags:        nil,
			}),
			tracer,
		)

		h.Use(hertztracing.ServerMiddleware(cfg), hertzSentinel.SentinelServerMiddleware(
		//// customize resource extractor if required
		//// method_path by default
		//hertzSentinel.WithServerResourceExtractor(func(c context.Context, ctx *app.RequestContext) string {
		//	return "server_test"
		//}),
		//// customize block fallback if required
		//// abort with status 429 by default
		//hertzSentinel.WithServerBlockFallback(func(c context.Context, ctx *app.RequestContext) {
		//	ctx.AbortWithStatusJSON(400, utils.H{
		//		"err":  "too many request; the quota used up",
		//		"code": 10222,
		//	})
		//}),
		))

		router.InitRouter(h)
		h.Spin()
	}()

	wg.Wait()

}
