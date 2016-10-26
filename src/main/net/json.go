package net

import (
	"encoding/json"
	"errors"
	"fmt"
	"main/pojo"
)

func PersonJson() (string, error) {
	p := pojo.Person{100, "lin", 23}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("parse to json error")
		err := errors.New("parse json error")
		return "", err
	}
	str := string(b)
	fmt.Println(str)
	return str, nil
}

func JsonPerson(jsonstr string) (p pojo.Person, err error) {
	bb := []byte(jsonstr)
	var ppp pojo.Person
	//必须要穿入实例化对象
	err1 := json.Unmarshal(bb, &ppp)
	if err1 != nil {
		fmt.Println("parse json to object error")
		errors.New("parse json to object error")
		return
	}
	fmt.Println(ppp)
	return ppp, nil
}
