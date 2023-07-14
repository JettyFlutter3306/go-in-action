package net_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/julienschmidt/httprouter"
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

	// 阻塞客户端端,不放服务终止
	for {
		// 接收服务端信息
		n, err := response.Body.Read(buffer)

		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			fmt.Println("读取完毕")
			res := string(buffer[:n]) // 转为切片再转为字符串
			fmt.Println(res)
			break
		}
	}
}

func TestHttpServer(t *testing.T) {
	address := "127.0.0.1:8888"
	content := "http://www.baidu.com 大傻逼哈哈哈"
	buffer := []byte(content)

	// 单独写回调函数,使用匿名函数,参考JQuery的Ajax
	http.HandleFunc("/go", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.RemoteAddr, "连接成功!")
		fmt.Println("method: ", request.Method)
		fmt.Println("url: ", request.URL.Path)
		fmt.Println("header: ", request.Header)
		fmt.Println("body: ", request.Body)

		// 响应
		_, _ = writer.Write(buffer)
	})

	http.HandleFunc("/params", func(writer http.ResponseWriter, request *http.Request) {
		username := request.FormValue("username")
		age := request.FormValue("age")
		address := request.FormValue("address")
		fmt.Println(username)
		fmt.Println(age)
		fmt.Println(address)

		_, _ = writer.Write([]byte("收到请求"))
	})

	err := http.ListenAndServe(address, nil)
	if err != nil {
		return
	}
}

func TestRequestParams(t *testing.T) {
	values := url.Values{}
	values.Set("username", "洛必达")
	values.Set("age", "18")
	values.Set("address", "暴风城")
	u, err := url.Parse("http://localhost:8888/params")
	if err != nil {
		fmt.Println(err)
	}

	response, err := http.PostForm(u.String(), values)
	if err != nil {
		fmt.Println(err)
	}

	bytes, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

func TestMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("hello mux")
	})
}

func TestHttpRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		writer.Write([]byte("hello httpRouter!!"))
	})
	router.GET("/index/:name", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "hello %s \n", params.ByName("name"))
	})
	http.ListenAndServe(":8081", router)
}
