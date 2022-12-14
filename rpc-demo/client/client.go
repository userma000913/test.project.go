package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	"test.project/rpc-demo"
)

func main() {

	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var res float64
	err = client.Call("DemoService.Div", rpc_demo.Arg{A: 10, B: 3}, &res)

	fmt.Println(res, err)
}
