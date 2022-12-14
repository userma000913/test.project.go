package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"test.project/rpc-demo"
)

func main() {

	err := rpc.Register(rpc_demo.DemoService{})
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Println(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
