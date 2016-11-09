package main

import (
	"fmt"
	my "main/net"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func main() {

	rpcService := &my.MyRpcService{}
	rpcService.My = &my.MyService{}

	rpc.Register(rpcService)
	rpc.HandleHTTP()

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:1234")

	listener, err := net.ListenTCP("tcp", addr)

	if err != nil {
		fmt.Println("hahaha ")
		return
	}

	go http.Serve(listener, nil)

	fmt.Println("rpc server started")

	time.Sleep(time.Duration(10000) * time.Second)

}
