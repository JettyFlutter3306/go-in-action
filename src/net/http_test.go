package net

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

/*
Go语言内置的net/http包提供了HTTP客户端和服务端的实现。
超文本传输协议（HTTP，HyperText Transfer Protocol)是互联网上应用最为广泛的一种网络传输协议，
所有的WWW文件都必须遵守这个标准。
设计HTTP最初的目的是为了提供一种发布和接收HTML页面的方法。
*/
func TestHttp(t *testing.T) {
	//基本的HTTP/HTTPS请求 Get、Head、Post和 PostForm 函数发出HTTP/HTTPS请求。
	var siteUrl = "https://element.eleme.cn/#/zh-CN/component/icon"

	response, err := http.Get(siteUrl)

	//resp, err := http.Post("http://5lmh.com/upload", "image/jpeg", &buf)

	//resp, err := http.PostForm("http://5lmh.com/form", url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		//handle err
		fmt.Println("read from resp.Body failed,err:", err)
		return
	}

	defer response.Body.Close() //用完response之后最后关闭回复的主体

	body, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))
}

/*
将上面的代码保存之后编译成可执行文件，执行之后就能在终端打印网站首页的内容了，
我们的浏览器其实就是一个发送和接收HTTP协议数据的客户端，
我们平时通过浏览器访问网页其实就是从网站的服务器接收HTTP数据，
然后浏览器会按照HTML、CSS等规则将网页渲染展示出来。
*/

func TestArgsGet(t *testing.T) {
	apiUrl := "http://127.0.0.1:9090/get"

	// URL param
	data := url.Values{}
	data.Set("name", "枯藤")
	data.Set("age", "18")

	u, err := url.ParseRequestURI(apiUrl)

	if err != nil {
		fmt.Printf("parse url requestUrl failed,err:%v\n", err)
	}

	u.RawQuery = data.Encode() // URL encode
	fmt.Println(u.String())

	resp, err := http.Get(u.String())

	if err != nil {
		fmt.Printf("post failed, err: %v\n", err)
		return
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("get resp failed,err: %v\n", err)
		return
	}

	fmt.Println(string(b))
}

//测试post
func TestPost(t *testing.T) {
	address := "http://127.0.0.1:9090/post"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=枯藤&age=18"
	// json

	contentType := "application/json"
	data := `{"name":"枯藤","age":18}`

	resp, err := http.Post(address, contentType, strings.NewReader(data))

	if err != nil {
		fmt.Printf("post failed, err: %v \n", err)
		return
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("get resp failed,err: %v \n", err)
		return
	}

	fmt.Println(string(b))
}
