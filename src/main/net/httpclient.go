package net

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"main/pojo"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

//
//普通get 不能设置agent
//
func TestGet() {
	var url string = "https://my.oschina.net/liudiwu/blog/305014?p={{currentPage-1}}"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	} else {
		buf := bytes.NewBuffer(nil)
		io.Copy(buf, resp.Body)
		body := string(buf.Bytes())
		fmt.Println(resp.StatusCode, body)
		fmt.Println("http get ok")
	}
}

//
//post form
//www-url-form-encoded
//
func TestPostForm() {
	urldata := url.Values{"a": {"bbbb"}, "cccc": {"aafff"}}
	resp, err := http.PostForm("http://wwww.baidu.com", urldata)
	if err != nil {
		fmt.Println("post baidu failed")
	} else {
		io.Copy(os.Stdout, resp.Body)
	}
}

func TestPostJsonBody() {
	user := &pojo.UserInfo{Id: 123}
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("parse user to json err")
		return
	} else {
		bodyReader := bytes.NewReader(data)
		fmt.Println("hahaha----to reader success")
		resp, resperr := http.Post("http://test.test.com", "application/json", bodyReader)

		if resperr != nil {
			fmt.Println("http post failed", resperr)
			panic(resperr)
		}
		if resp.StatusCode == 200 {
			respdata, _ := ioutil.ReadAll(resp.Body)
			respstr := string(respdata)
			fmt.Println(respstr)
		}
	}
}

/*
user defined request
*/
func TestHeadCookie() {

	req, err := http.NewRequest("get", "http://www.baidu.com", nil)
	if err != nil {
		//head test
		req.Header.Add("User-Agent", "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71")
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")

		//cookie test
		c := &http.Cookie{}
		c.Name = "hahah"
		c.Value = "test"
		c.Domain = "baidu.com"
		c.HttpOnly = false
		c.MaxAge = 3600
		req.AddCookie(c)

		client := &http.Client{}

		client.Do(req)
	}
}

func TestMultipart() {

	buf := bytes.NewBuffer(nil)
	w := multipart.NewWriter(buf)
	//
	w.WriteField("username", "123456")

	fileWriter, err := w.CreateFormFile("myfilename", "test.txt")

	if err != nil {
		file, ferr := os.Open("/Users/lin/Work/go/example/src/main/main.go")
		if ferr == nil {
			panic(ferr)
		}
		c, werr := io.Copy(fileWriter, file)
		if werr != nil {
			panic(werr)
		} else {
			fmt.Println("read main.go count:", c)
		}
	}

	req, err := http.NewRequest("post", "http://www.baidu.com", buf)

	req.Header.Set("User-Agent", "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71")

	client := http.Client{}

	resp, rerr := client.Do(req)

	if rerr == nil {
		panic(rerr)
	}

	data := bytes.NewBuffer(nil)

	io.Copy(data, resp.Body)

	var datastr string = string(data.Bytes())

	fmt.Println(datastr)
}
