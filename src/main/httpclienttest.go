package main

import (
	"fmt"
	"main/net"
)

func TestGet() {
	net.TestGet()
}

func TestTaobao() {
	resp := net.TestTaobaoIP("115.236.39.210")
	fmt.Println(resp)
}

func main() {
	TestTaobao()

}
