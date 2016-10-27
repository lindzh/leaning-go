package net

import (
	"fmt"
	"html/template"
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
