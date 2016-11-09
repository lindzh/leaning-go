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
	"strings"
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

func TestTaobaoIP(ip string) string {
	var requesUrl string = "http://ip.taobao.com/service/getIpInfo.php?ip=" + ip
	resp, err := http.Get(requesUrl)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err2)
		return ""
	}

	return string(data)
}

func TestPostData() {
	proxy := func(r *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:8888")
	}

	transport := &http.Transport{Proxy: proxy}

	var requestUrl = "http://gateway.aidaojia.com/gateway"

	var stringBody = `{"a":"test"}`

	// resp, _ := http.Post(requestUrl, "application/json", body)

	req, _ := http.NewRequest("POST", requestUrl, strings.NewReader(stringBody))

	req.Header.Add("Content-Type", "application/json")

	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36")

	req.Header.Add("Accept", "*/*")

	req.Header.Add("Accept-Encoding", "gzip, deflate")

	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6")

	req.Header.Add("Cache-Control", "no-cache")

	client := &http.Client{Transport: transport}

	// client := http.Client{}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	fmt.Println(err)

	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))

}

func TestPostDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.01happy.com/demo/accept.php", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
