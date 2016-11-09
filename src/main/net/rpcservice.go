package net

import (
	"errors"
	"main/pojo"
)

type MyRpcService struct {
	My *MyService
}

func (service *MyRpcService) SimpleRpc(person *pojo.Person, user *pojo.UserInfo) error {
	if person == nil {
		return errors.New("person is null")
	}

	u := service.My.SayHello(person)
	user.Age = u.Age
	user.Email = u.Email
	user.Id = u.Id
	user.Links = u.Links
	user.Mobile = u.Mobile
	user.Name = u.Name
	return nil
}
