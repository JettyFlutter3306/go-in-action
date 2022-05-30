package net_test

import (
	"fmt"
	"net"
	"testing"
)

func TestTcpClient(t *testing.T) {
	fmt.Println("客户端启动...")
	conn, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		fmt.Println("客户端连接失败! err: ", err)
		return
	}
	fmt.Println("连接成功! conn: ", conn)

	str := "Hello World!!!"

	n, err := conn.Write([]byte(str))
	if err != nil {
		fmt.Println("连接失败! err: ", err)
	}
	fmt.Printf("终端发送数据成功!一共发送了 %d 字节的数据\n", n)
}

func TestTcpServer(t *testing.T) {
	fmt.Println("服务器启动了...")
	listen, err := net.Listen("tcp", "127.0.0.1:8080")

	if err != nil {
		fmt.Println("监听失败! err: ", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("客户端等待失败! err: ", err)
		} else {
			fmt.Printf("等待连接成功! conn = %v, IP & Port: %v \n", conn, conn.RemoteAddr().String())
		}

		// 准备一个协程
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		fmt.Println(string(buffer[0:n]))
	}
}
