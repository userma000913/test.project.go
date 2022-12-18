package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	hertzSentinel "github.com/hertz-contrib/opensergo/sentinel/adapter"
	"github.com/hertz-contrib/registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"hertz_demo/conf"
	"hertz_demo/controller"
	"hertz_demo/router"
	"hertz_demo/service"
	"log"
)

func main() {

	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848),
	}

	cc := constant.ClientConfig{
		NamespaceId:         "public",
		Username:            "nacos",
		Password:            "nacos",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	n := conf.InitNacos(sc, cc)
	c := conf.InitConfigWithNacos(n)
	if c == nil {
		panic("config is nil")
	}
	s := service.New(c)
	controller.Init(s)
	addr := fmt.Sprintf("127.0.0.1:%d", c.Port)

	// nacos注册中心客户端
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

	// 服务注册
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

	// Tracing & Sentinel
	h.Use(hertztracing.ServerMiddleware(cfg), hertzSentinel.SentinelServerMiddleware())

	router.InitRouter(h)
	h.Spin()

}
