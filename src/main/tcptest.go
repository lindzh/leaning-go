package main

import (
	"fmt"
	tcp "main/net"
	"time"
)

func main() {
	tcp.ServerStart()
	go tcp.ClientStart()
	fmt.Println("client server started")
	time.Sleep(time.Duration(10) * time.Second)
}
