package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"io"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	hertztracer "github.com/hertz-contrib/tracer/hertz"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

/*
export JAEGER_DISABLED=false
export JAEGER_SAMPLER_TYPE="const"
export JAEGER_SAMPLER_PARAM=1
export JAEGER_REPORTER_LOG_SPANS=true
export JAEGER_AGENT_HOST="127.0.0.1"
export JAEGER_AGENT_PORT=6831
*/

// InitTracer Initialize and create tracer
func InitTracer(serviceName string) (opentracing.Tracer, io.Closer) {
	cfg, _ := jaegercfg.FromEnv()
	cfg.ServiceName = serviceName
	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot initialization Jaeger: %v\n", err))
	}
	// opentracing.InitGlobalTracer(tracer)
	return tracer, closer
}

func main() {
	ht, hc := InitTracer("hertz-server")
	//kt, kc := InitTracer("kitex-client")
	defer hc.Close()
	//defer kc.Close()

	// kitex-client configure tracer
	//client, err := echo.NewClient("echo",
	//	kclient.WithHostPorts("0.0.0.0:5555"),
	//	kclient.WithSuite(kopentracing.NewClientSuite(kt, func(c context.Context) string {
	//		endpoint := rpcinfo.GetRPCInfo(c).From()
	//		return endpoint.ServiceName() + "::" + endpoint.Method()
	//	})))
	//if err != nil {
	//	panic(err)
	//}

	// hertz-server configure tracer
	h := server.Default(server.WithTracer(hertztracer.NewTracer(ht, func(c *app.RequestContext) string {
		return "test.hertz.server" + "::" + c.FullPath()
	})))

	// Register and use tracer middleware.
	// This middleware is simple demo.
	// You can refer to the example to implement a tracer middleware yourself to get the metrics you want.
	h.Use(hertztracer.ServerCtx())

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		type PingReq struct {
			Name string `query:"name"`
		}
		var hertzReq PingReq
		err := ctx.BindAndValidate(&hertzReq)
		if err != nil {
			hlog.Errorf(err.Error())
			return
		}

		//KitexReq := &api.Request{Message: hertzReq.Name}
		//resp, err := client.Echo(c, KitexReq)
		//if err != nil {
		//	hlog.Errorf(err.Error())
		//}
		ctx.JSON(consts.StatusOK, hertzReq)
	})

	h.Spin()
}
