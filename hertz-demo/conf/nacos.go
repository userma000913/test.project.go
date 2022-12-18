package conf

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

type nacosClient struct {
	Sc []constant.ServerConfig
	Cc constant.ClientConfig
}

func NewNacosClient() *nacosClient {
	return &nacosClient{}

}

func (n *nacosClient) InitNacosClient() {
	n.Sc = []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848),
	}

	n.Cc = *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithUsername("nacos"),
		constant.WithPassword("nacos"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

}

// nacos做配置中心
func (n *nacosClient) GetNacosConfigClient() config_client.IConfigClient {
	// 创建nacos config客户端
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &n.Cc,
			ServerConfigs: n.Sc,
		},
	)

	if err != nil {
		log.Println("nacos client创建失败 。。。", err.Error())
		return nil
	}

	//content, err := client.GetConfig(vo.ConfigParam{
	//	DataId: "hertz_demo",
	//	Group:  "DEFAULT_GROUP"})
	//
	//if err != nil {
	//	log.Println("配置获取失败 。。。", err.Error())
	//}

	return client

}

func (n *nacosClient) GetNacosConfigContent() (string, error) {
	content, err := n.GetNacosConfigClient().
		GetConfig(vo.ConfigParam{DataId: "hertz_demo", Group: "DEFAULT_GROUP"})
	if err != nil {
		log.Println("配置获取失败 。。。", err.Error())
		return "", err
	}
	return content, nil
}
