package main

import (
	"fmt"
	"main/pojo"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("dail http failed")
	}

	person := &pojo.Person{}
	person.Id = 12345
	person.Name = "lindeshzi"
	person.Age = 21

	user := &pojo.UserInfo{}

	cerr := client.Call("MyRpcService.SimpleRpc", person, user)

	if cerr != nil {
		fmt.Println(cerr)
	}

	fmt.Println(*user)
}
