package impl

import (
	"context"
	"errors"
	"fmt"
	"grpc_demo/proto/say_hello"
)

type SayHelloServiceImpl struct {
	say_hello.UnimplementedSayHelloServiceServer
}

// SayHello 实现接口
func (SayHelloServiceImpl) SayHello(ctx context.Context, req *say_hello.SayHelloRequest) (resp *say_hello.SayHelloResponse, err error) {
	resp = &say_hello.SayHelloResponse{}
	if req.Request == "" {
		return nil, errors.New("req Request is invalid")
	}
	resp.Response = fmt.Sprintf("hello %s ~", req.GetRequest())
	return
}
