package net_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestHttpClient(t *testing.T) {
	address := "http://127.0.0.1:8888/go"

	response, err := http.Get(address)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	fmt.Println(response.Status)
	fmt.Println(response.Header)

	buffer := make([]byte, 1024)

	//阻塞客户端端,不放服务终止
	for true {
		//接收服务端信息
		n, err := response.Body.Read(buffer)

		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			fmt.Println("读取完毕")
			res := string(buffer[:n]) //转为切片再转为字符串
			fmt.Println(res)
			break
		}
	}
}

func TestHttpServer(t *testing.T) {
	address := "127.0.0.1:8888"

	content := "http://www.baidu.com 大傻逼哈哈哈"
	buffer := []byte(content)

	//单独写回调函数,使用匿名函数,参考JQuery的Ajax
	http.HandleFunc("/go", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.RemoteAddr, "连接成功!")

		//请求方式
		fmt.Println("method: ", request.Method)
		fmt.Println("url: ", request.URL.Path)
		fmt.Println("header: ", request.Header)
		fmt.Println("body: ", request.Body)

		//回复
		_, _ = writer.Write(buffer)
	})

	http.ListenAndServe(address, nil)
}
