package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/registry/nacos"
	"hertz_demo/config"
	"hertz_demo/dao"
	"hertz_demo/router"
	"log"
)

func main() {

	c := config.InitConfig()
	dao.InitMysql()
	addr := fmt.Sprintf("127.0.0.1:%d", c.Port)

	// 服务注册
	r, err := nacos.NewDefaultNacosRegistry()
	if err != nil {
		log.Fatal(err)
		return
	}
	h := server.Default(
		server.WithHostPorts(addr),
		server.WithRegistry(r, &registry.Info{
			ServiceName: c.Name,
			Addr:        utils.NewNetAddr("tcp", addr),
			Weight:      10,
			Tags:        nil,
		}),
	)
	router.InitRouter(h)
	h.Spin()

}
