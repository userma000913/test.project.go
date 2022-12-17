package main

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

// Nacos作为注册中心
func main() {
	initRegisterNacos()
	for {

	}
}

func initRegisterNacos() {
	sc := []constant.ServerConfig{
		{
			IpAddr: "localhost",
			Port:   8848,
		},
	}

	// nacos client conf
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(""), //当namespace是public时，此处填空字符串。
		constant.WithUsername("nacos"),
		constant.WithPassword("nacos"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
	)
	// 创建服务发现客户端的另一种方式 (推荐)
	namingClient, _ := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "localhost",
		Port:        8081,
		ServiceName: "nacos-demo",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
	})
}
