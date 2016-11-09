package net

import (
	"fmt"
	"main/pojo"
)

type MyService struct {
}

func (service *MyService) SayHello(person *pojo.Person) *pojo.UserInfo {
	fmt.Println("say Hello person", *person)

	user := &pojo.UserInfo{}

	user.Age = 11
	user.Email = "linsony0@163.com"
	user.Id = 100
	user.Mobile = "15868884065"
	user.Name = "rpclin"
	user.Links = []string{"hahaha", "123456"}
	return user
}
