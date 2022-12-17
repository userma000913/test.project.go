package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	hertzSentinel "github.com/hertz-contrib/opensergo/sentinel/adapter"

	"github.com/cloudwego/hertz/pkg/common/utils"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/registry/nacos"
	"hertz_demo/conf"
	"hertz_demo/controller"
	"hertz_demo/router"
	"hertz_demo/service"
	"log"
)

func main() {

	c := conf.InitConfig()
	s := service.New(c)
	controller.Init(s)
	addr := fmt.Sprintf("127.0.0.1:%d", c.Port)

	// 服务注册
	r, err := nacos.NewDefaultNacosRegistry()
	if err != nil {
		log.Fatal(err)
		return
	}

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

}
