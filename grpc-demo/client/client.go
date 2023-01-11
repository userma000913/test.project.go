package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo/proto/say_hello"
)

var serverPort = 8888

func main() {
	conn, err := grpc.Dial(fmt.Sprintf(":%d", serverPort))
	if err != nil {

	}
	defer conn.Close()

	client := say_hello.NewSayHelloServiceClient(conn)

	req := &say_hello.SayHelloRequest{
		Request: "Tom",
	}
	resp, err := client.SayHello(context.Background(), req)
	if err != nil {
	}
	fmt.Println(resp.GetResponse())

}
