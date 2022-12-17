package proxy

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/registry/nacos"
	"log"
)

type HTTP struct {
	ServiceName string
}

func InitHTTP(serviceName string) *HTTP {
	return &HTTP{ServiceName: serviceName}
}

// Resolver 服务发现
func (h *HTTP) Resolver() *client.Client {
	c, err := client.NewClient()
	// OpenTelemetry
	c.Use(hertztracing.ClientMiddleware())

	if err != nil {
		panic(err)
	}
	r, err := nacos.NewDefaultNacosResolver()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	c.Use(sd.Discovery(r))
	return c
}

func (h *HTTP) JsonPost(ctx context.Context, url string, data []byte) (int, []byte, error) {
	url = fmt.Sprintf("http://%s/%s", h.ServiceName, url)
	req := protocol.AcquireRequest()
	req.SetOptions(config.WithSD(true))
	req.SetMethod(consts.MethodPost)
	req.SetRequestURI(url)
	req.SetBody(data)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	resp := protocol.AcquireResponse()
	err := client.Do(ctx, req, resp)
	if err != nil {
		hlog.Fatal(err)
		return resp.StatusCode(), nil, err
	}
	return resp.StatusCode(), resp.Body(), nil

}

func (h *HTTP) FormPost(ctx context.Context, url string, data map[string]string) (int, []byte, error) {
	url = fmt.Sprintf("http://%s/%s", h.ServiceName, url)
	req := protocol.AcquireRequest()
	req.SetOptions(config.WithSD(true))
	req.SetMethod(consts.MethodPost)
	req.SetRequestURI(url)
	req.Header.SetContentTypeBytes([]byte("application/x-www-form-urlencoded"))
	req.SetFormData(data)
	resp := protocol.AcquireResponse()
	err := client.Do(ctx, req, resp)
	if err != nil {
		hlog.Fatal(err)
		return resp.StatusCode(), nil, err
	}
	return resp.StatusCode(), resp.Body(), nil
}

func (h *HTTP) Get(ctx context.Context, url string) {
	url = fmt.Sprintf("http://%s/%s", h.ServiceName, url)
	status, body, err := h.Resolver().Get(ctx, nil, url, config.WithSD(true))
	if err != nil {
		hlog.Fatal(err)
	}
	hlog.Infof("code=%d,body=%s\n", status, string(body))
}
