package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
)

// 配置中心
func main() {

}

func initNacos() {
	// nacos server config
	sc := []constant.ServerConfig{
		{
			IpAddr:   "nacos.apizones.com",
			Port:     443,
			GrpcPort: 9848,
			Scheme:   "https",
		},
	}

	// nacos client config
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId("bd71943c-a8df-4e8e-8d27-eda672f3550a"),
		constant.WithUsername("dev"),
		constant.WithPassword("e77989ed21758e78331b20e477fc5582"),
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
		return
	}

	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "msg_config_test",
		Group:  "DEFAULT_GROUP"})

	if err != nil {
		log.Println("配置获取失败 。。。", err.Error())
	}

	// 配置信息
	fmt.Printf("%s", content)
}
