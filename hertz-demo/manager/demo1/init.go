package demo1

import (
	"hertz_demo/conf"
	"hertz_demo/proxy"
)

type Manager struct {
	c          *conf.AppConfig
	httpClient *proxy.HTTP
	Url1       string
}

func NewManager(c *conf.AppConfig) *Manager {
	return &Manager{
		c:          c,
		httpClient: proxy.InitHTTP("demo1"),
		Url1:       "/ping",
	}
}
