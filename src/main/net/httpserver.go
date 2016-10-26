package net

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func hello(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "welcome to the go web site ", time.Now())
}

// get post
func login(response http.ResponseWriter, request *http.Request) {
	//parse form
	err := request.ParseForm()
	if err != nil {
		fmt.Println("http request parse form failed")
		dat := []byte("http server error")
		response.Write(dat)
	} else {
		//get value
		username := request.FormValue("username")
		password := request.FormValue("password")

		result := make(map[string]interface{})
		result["user"] = username
		result["hahah"] = 123
		result["pass"] = password
		buf, er := json.Marshal(result)
		if er != nil {
			response.Write([]byte("http server error"))
		} else {
			response.Write(buf)
		}

	}

}

func head(response http.ResponseWriter, request *http.Request) {

}

func body(response http.ResponseWriter, request *http.Request) {

}

func cookie(response http.ResponseWriter, request *http.Request) {

}

func defaultHandler(response http.ResponseWriter, request *http.Request) {

}

func StartHttpServer() {
	var addr = "127.0.0.1:8088"
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/head", head)
	http.HandleFunc("/body", body)
	http.HandleFunc("cookie", cookie)
	fmt.Println("http server starting")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("start http server failed")
	} else {
		fmt.Println("start http server success")
	}

}
