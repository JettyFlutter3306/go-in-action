package main

import (
	"fmt"
	"net/http"
)

func UserResp(resp http.ResponseWriter, req *http.Request) {
	fmt.Printf("请求方法: %s \n", req.Method)
	fmt.Printf("浏览器发送请求文件路径: %s \n", req.URL)
	fmt.Printf("请求头: %s \n", req.Header)
	fmt.Printf("请求包体: %s \n", req.Body)
	fmt.Printf("客户端网络地址: %s \n", req.RemoteAddr)
	fmt.Printf("客户端Agent: %s \n", req.UserAgent())

	/**
	  给客户端回复数据
	*/
	resp.Write([]byte("User response"))
}

func IndexResp(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Index response"))
}

func main() {
	/**
	  为不同的请求注册不同的函数
	*/
	http.HandleFunc("/user", UserResp)
	http.HandleFunc("/index", IndexResp)

	//开启服务器，监听客户端的请求
	http.ListenAndServe("127.0.0.1:8888", nil)

}
