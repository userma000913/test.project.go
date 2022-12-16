package initialization

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
)

// nacos做配置中心
func InitNacosConfig() string {
	// nacos server config
	sc := []constant.ServerConfig{
		{
			IpAddr: "localhost",
			Port:   8848,
			Scheme: "http",
		},
	}

	// nacos client config
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithUsername("nacos"),
		constant.WithPassword("nacos"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)
	// 创建nacos config客户端
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		log.Println("nacos client创建失败 。。。", err.Error())
		return ""
	}

	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "hertz_demo",
		Group:  "DEFAULT_GROUP"})

	if err != nil {
		log.Println("配置获取失败 。。。", err.Error())
	}

	return content

}
