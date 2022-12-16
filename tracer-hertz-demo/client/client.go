package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"io"
	"time"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertztracer "github.com/hertz-contrib/tracer/hertz"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func main() {
	closer := InitJaeger("hertz-client")
	defer closer.Close()
	c, _ := client.NewClient()

	// Register and use client tracer middleware.
	// This middleware is simple demo.
	// You can refer to the example to implement a tracer middleware yourself to get the metrics you want.
	c.Use(hertztracer.ClientTraceMW, hertztracer.ClientCtx)
	for {
		_, b, err := c.Get(context.Background(), nil, "http://localhost:8888/ping?name=hertz")
		if err != nil {
			hlog.Errorf(err.Error())
		}
		hlog.Infof(string(b))
		time.Sleep(time.Second)
	}
}

// InitJaeger ...
func InitJaeger(service string) io.Closer {
	cfg, _ := jaegercfg.FromEnv()
	cfg.ServiceName = service
	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot initialization Jaeger: %v\n", err))
	}
	opentracing.InitGlobalTracer(tracer)
	return closer
}
