package net

import (
	"fmt"
	"html/template"
	"main/pojo"
	"os"
)

var content = ""

/*
简单模版测试s
*/
func TestHello() {
	var src = "hello this is {{.name}}"
	t := template.New("hello")
	tpl, err := t.Parse(src)
	if err != nil {
		fmt.Println(err)
	}
	data := make(map[string]string, 10)
	data["name"] = "hahah"
	tpl.Execute(os.Stdout, data)
}

/*
测试for循环
*/
func TestObjectAndList() {
	var links []string = []string{"123", "hahah", "tat vv", "this"}
	user := pojo.UserInfo{}
	user.Id = 123
	user.Age = 21
	user.Email = "linsony@gmail.com"
	user.Mobile = "15868884065"
	user.Name = "lindezhi"
	user.Links = links
	data := make(map[string]interface{})
	data["user"] = user
	// user := pojo.UserInfo(123, "lin", 21, "15868884061", "linsony@gmail.com", links)
	var src string = `hello, my name is {{.user.Name}}, and my id is {{.user.Id}} my mobile is {{.user.Mobile}} and my address is {{range .user.Links}}{{.}},{{end}}`
	t := template.New("haha")
	tpl, err := t.Parse(src)
	if err != nil {
		fmt.Println(err)
	}
	tpl.Execute(os.Stdout, data)
}

/*
测试 if else
*/
func TestIfElse() {
	tpl := template.New("hahah")
	var src string = `hahah today is {{if .friday}} friday. {{else}} hahah{{end}}`
	tt, _ := tpl.Parse(src)
	data := make(map[string]interface{})
	data["friday"] = true
	tt.Execute(os.Stdout, data)

}
