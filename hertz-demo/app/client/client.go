package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/registry/nacos"
	"log"
)

func main() {
	client, err := client.NewClient()
	if err != nil {
		panic(err)
	}
	r, err := nacos.NewDefaultNacosResolver()
	if err != nil {
		log.Fatal(err)
		return
	}
	client.Use(sd.Discovery(r))

	for i := 0; i < 10; i++ {

		//r := struct {
		//	Name string `json:"name" form:"name"`
		//	Age  int    `json:"age" form:"age"`
		//}{
		//	Name: fmt.Sprintf("test%d", i),
		//	Age:  i,
		//}
		//b, err := json.Marshal(r)
		if err != nil {
			fmt.Println("Marshal err")
			continue
		}
		//
		//req := &protocol.Args{}
		//req.AppendBytes(b)
		//
		//status, body, err := client.Post(context.Background(), []byte{}, "http://hertz.test.demo/ping", req, config.WithSD(true))
		//if err != nil {
		//	hlog.Fatal(err)
		//}
		//hlog.Infof("code=%d,body=%s\n", status, string(body))

		req := protocol.AcquireRequest()
		req.SetOptions(config.WithSD(true))
		req.SetMethod(consts.MethodPost)
		req.SetRequestURI("http://demo1/ping")
		type Test struct {
			A int `json:"a"`
			B int `json:"b"`
		}
		t := Test{A: 11, B: 22}
		b, err := json.Marshal(t)
		if err != nil {

		}
		req.SetBody(b)
		//req.SetFormData(map[string]string{
		//	"name": "tom",
		//	"age":  "11",
		//})
		req.Header.SetContentTypeBytes([]byte("application/json"))
		resp := protocol.AcquireResponse()
		err = client.Do(context.Background(), req, resp)
		if err != nil {
			hlog.Fatal(err)
		}
		hlog.Infof("code=%d,body=%s", resp.StatusCode(), string(resp.Body()))

		//status, body, err := client.Get(context.Background(), nil, "http://hertz.test.demo/t", config.WithSD(true))
		//if err != nil {
		//	hlog.Fatal(err)
		//}
		//hlog.Infof("code=%d,body=%s\n", status, string(body))
	}
}
