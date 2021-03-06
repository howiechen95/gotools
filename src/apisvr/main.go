package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func WebServerBase() {
	fmt.Println("This is webserver base!")

	//第一个参数为客户端发起http请求时的接口名，第二个参数是一个func，负责处理这个请求。
	http.HandleFunc("/test", test)

	http.HandleFunc("/test/chatroom", testChatRoom)

	//服务器要监听的主机地址和端口号
	err := http.ListenAndServe("127.0.0.1:8080", nil)

	if err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
	}
}

func test(w http.ResponseWriter, req *http.Request) {
	fmt.Println("loginTask is running...")

	//模拟延时
	//time.Sleep(time.Second * 2)

	//获取客户端通过GET/POST方式传递的参数
	req.ParseForm()
	param_userName, found1 := req.Form["userName"]
	param_password, found2 := req.Form["password"]

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", b)

	if !(found1 && found2) {
		fmt.Fprint(w, "请勿非法访问")
		return
	}

	result := NewBaseJsonBean()
	userName := param_userName[0]
	password := param_password[0]

	s := "userName:" + userName + ",password:" + password
	fmt.Println(s)

	if userName == "zhangsan" && password == "123456" {
		result.Code = 100
		result.Message = "登录成功"
	} else {
		result.Code = 101
		result.Message = "用户名或密码不正确"
	}

	//向客户端返回JSON数据
	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))
}

func testChatRoom(w http.ResponseWriter, req *http.Request) {
	result := NewBaseJsonBean()

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("testChatRoom")
	fmt.Printf("%s", b)

	result.Code = 100
	result.Message = "请求成功 /test/chatroom"
	//向客户端返回JSON数据
	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))
}

type BaseJsonBean struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewBaseJsonBean() *BaseJsonBean {
	return &BaseJsonBean{}
}

func main() {
	WebServerBase()
}
