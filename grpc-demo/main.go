package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo/impl"
	"grpc_demo/proto/say_hello"
	"net"
)

var port = 8888

func main() {
	server := grpc.NewServer()
	say_hello.RegisterSayHelloServiceServer(server, new(impl.SayHelloServiceImpl))

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
